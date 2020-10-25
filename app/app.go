package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	marketdata "github.com/isavinof/pricer/market_data/market_data_provider"
	productparser "github.com/isavinof/pricer/product_parser/csv_product_parser"
	"github.com/isavinof/pricer/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/isavinof/pricer/config"
	"github.com/isavinof/pricer/log"
	marketdatastore "github.com/isavinof/pricer/market_data/market_data_store"
	"github.com/isavinof/pricer/network"
	pricelist "github.com/isavinof/pricer/price-list"
	pricelistserver "github.com/isavinof/pricer/price-list-server"
	systemsignals "github.com/isavinof/pricer/system_signals"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type App struct {
	Logger          log.Logger
	config          config.Config
	store           pricelistserver.MarketDataStore
	provider        pricelistserver.MarketDataProvider
	grpcServer      *grpc.Server
	mongoClient     *mongo.Client
	listener        net.Listener
	signalsListener systemsignals.SignalListener
	ctxCancel       func()
	ctx             context.Context
}

func NewApp(envs []string) (*App, error) {
	logger := log.NewLogger()
	app := &App{Logger: logger}

	logger.Info("start parse config:", envs)
	err := app.parseConfig(envs)
	if err != nil {
		return nil, err
	}

	app.ctx, app.ctxCancel = context.WithCancel(log.ToContext(context.Background(), app.Logger))

	logger.Info("start initialize store")
	err = app.initializeStore(app.ctx, app.config.Mongo)
	if err != nil {
		return nil, err
	}

	logger.Info("start initialize provider")
	err = app.initMarketDataProvider(app.ctx)
	if err != nil {
		return nil, err
	}

	logger.Info("start initialize network")
	err = app.initNetwork(app.ctx, app.config.GRPC)
	if err != nil {
		return nil, err
	}

	logger.Info("start initialize grpc")
	err = app.initGRPC(app.ctx, app.config.GRPC)
	if err != nil {
		return nil, err
	}

	logger.Info("start initialize signals")
	err = app.initSignals(app.ctx)
	if err != nil {
		return nil, err
	}

	logger.Info("register server")
	server := pricelistserver.NewPriceListServer(app.config.ServerConfig, app.Logger, app.provider, app.store, utils.RunWithTimeout)
	pricelist.RegisterPriceListServer(app.grpcServer, server)

	return app, nil
}

func (app *App) Run() error {
	app.Logger.Info("start listen signals")
	go app.signalsListener.ListenBlocked(app.ctx)

	app.Logger.Info("start db pinger")
	go app.RunPingerBlocked(app.ctx)

	app.Logger.Info("start test http")
	go app.StartHTTP(app.ctx)

	host, _ := os.Hostname()
	app.Logger.Infof("start listen grpc on: %v%v", host, app.config.GRPC.Port)
	err := app.grpcServer.Serve(app.listener)
	app.Logger.Info("listen grpc done")

	if err != nil {
		return err
	}

	return nil
}

func (app *App) RunPingerBlocked(ctx context.Context) {
	for {
		t := time.NewTimer(30 * time.Second)
		select {
		case <-t.C:
			rf := readpref.ReadPref{}
			err := app.mongoClient.Ping(context.Background(), &rf)
			if err != nil {
				app.grpcServer.GracefulStop()
				app.ctxCancel()
				app.Logger.Error("mongodb unreachable")
				return
			}
			app.Logger.Infof("mongodb health: %v", rf)
		case <-ctx.Done():
			app.Logger.Infof("context done. Stop ping mongodb")
			return
		}
	}
}

func (app *App) StartHTTP(ctx context.Context) {
	http.HandleFunc("/alive", func(writer http.ResponseWriter, request *http.Request) {
		rp := readpref.ReadPref{}
		err := app.mongoClient.Ping(request.Context(), &rp)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(rp)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Write(data)
		writer.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/products", func(writer http.ResponseWriter, request *http.Request) {
		fileName := fmt.Sprintf("%v_%v.csv", app.config.HTTP.FilePrefix, app.config.HTTP.RequestCount)
		app.Logger.Info("Serve file:", fileName)

		http.ServeFile(writer, request, fileName)
		// yes, it's local and unsafe. But this handler only for test
		app.config.HTTP.RequestCount++
		if app.config.HTTP.RequestCount > 5 {
			app.config.HTTP.RequestCount = 0
		}
	})
	addr := fmt.Sprintf(":%v", app.config.HTTP.Port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		app.Logger.Error("error to listen http on ", addr)
	}
}

func (app *App) parseConfig(envs []string) (err error) {
	app.Logger.Info("start parse config")
	app.config, err = config.ParseFromEnvs(envs)
	if err != nil {
		return errors.Wrap(err, "parse config")
	}

	app.Logger.Infof("config parsed:%v", app.config)
	return nil
}

func (app *App) initializeStore(ctx context.Context, mongoConfig config.MongoConfig) error {
	err := utils.RunWithTimeout(ctx, mongoConfig.ConnectTimeout, func(ctx context.Context) (err error) {
		app.Logger.Info("start create connection to mongo")
		app.mongoClient, err = marketdatastore.NewMongoConnection(ctx, mongoConfig)
		if err != nil {
			return err
		}
		err = app.mongoClient.Ping(ctx, nil)
		if err != nil {
			return err
		}
		app.Logger.Info("connected")
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "connect to db")
	}

	app.Logger.Info("start create index")
	store, err := marketdatastore.NewMongoStore(ctx, app.mongoClient.Database(mongoConfig.DB).Collection(mongoConfig.Collection))
	if err != nil {
		return errors.Wrapf(err, "initialize mongo error. %v", mongoConfig)
	}
	app.Logger.Info("store inited")
	app.store = store
	return nil
}

func (app *App) initMarketDataProvider(ctx context.Context) error {
	app.provider = marketdata.NewRestProvider(productparser.NewCsvParser())
	return nil
}

func (app *App) initGRPC(ctx context.Context, grpcConfig config.GRPCConfig) error {
	app.grpcServer = grpc.NewServer(grpcConfig.Options...)

	return nil
}

func (app *App) initNetwork(ctx context.Context, grpcConfig config.GRPCConfig) (err error) {
	addr := fmt.Sprintf(":%v", grpcConfig.Port)
	app.listener, err = network.NewTCPListener(ctx, network.ListenerOptions{HostPort: addr})
	if err != nil {
		return errors.Wrapf(err, "bind to the port:%v", addr)
	}

	return
}

func (app *App) initSignals(ctx context.Context) (err error) {
	app.signalsListener = systemsignals.NewSignalListener()
	app.signalsListener.SubscribeToShutdownSignals(func() {
		app.grpcServer.GracefulStop()
	})

	return nil
}
