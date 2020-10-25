package pricelistserver

import (
	"context"
	"time"

	"github.com/isavinof/pricer/config"

	"github.com/pkg/errors"

	"github.com/isavinof/pricer/log"

	pricelist "github.com/isavinof/pricer/price-list"
	"github.com/isavinof/pricer/types"
)

//go:generate mockgen -source=price_list_server.go -destination=price_list_server_mock.go -package=pricelistserver MarketDataProvider, MarketDataStore

// PriceListServer
type PriceListServer struct {
	provider       MarketDataProvider
	store          MarketDataStore
	runWithTimeout RunWithTimeout
	logger         log.Logger
	config         config.ServerConfig
}

// NewPriceListServer initialize PriceListServer object without any other actions
func NewPriceListServer(config config.ServerConfig, logger log.Logger, provider MarketDataProvider, store MarketDataStore, runWithTimeout RunWithTimeout) *PriceListServer {
	return &PriceListServer{
		provider:       provider,
		store:          store,
		logger:         logger,
		runWithTimeout: runWithTimeout,
		config:         config,
	}
}

// MarketDataProvider interface to request data from url
type MarketDataProvider interface {
	Fetch(ctx context.Context, url string) ([]types.ProductPrice, error)
}

// MarketDataStore interface to communicate with price store
type MarketDataStore interface {
	Save(ctx context.Context, prices []types.ProductPrice) error
	Get(ctx context.Context, sortType types.SortingType, directionType types.SortDirectionType, limit int64, offset int64) (prices []types.ProductPriceExtended, err error)
}

// RunWithTimeout function to run passed function with timeout
// Make tests easier
type RunWithTimeout = func(ctx context.Context, timeout time.Duration, f func(ctx context.Context) error) error

// UpdatePrices request product prices from external URL store it to PriceStore and return as response
// errors from this method returning as is.
// If we plan to use this server for external clients errors should be more clear
func (pls *PriceListServer) UpdatePrices(ctx context.Context, plRequest *pricelist.UpdatePriceListRequest) (*pricelist.UpdatePriceListResponse, error) {
	pls.logger.Info("receive update price request")
	ctx = log.ToContext(ctx, pls.logger)

	pls.logger.Info("start fetch prices")
	prices, err := pls.fetchPrices(ctx, plRequest.Url)
	if err != nil {
		return nil, err
	}

	pls.logger.Info("prices fetched. Size:%v", len(prices))
	if len(prices) > 0 {
		pls.logger.Info("start store prices")
		err = pls.storePrices(ctx, prices)
		if err != nil {
			return nil, err
		}
	}

	response := pricelist.UpdatePriceListResponse{Products: make([]*pricelist.ProductPrice, 0, len(prices))}
	for _, price := range prices {
		response.Products = append(response.Products, &pricelist.ProductPrice{
			ProductName:       price.ProductName,
			ProductPriceCents: price.ProductPriceCents,
		})
	}
	return &response, nil
}

// Get product prices from store
func (pls *PriceListServer) GetProductPrices(ctx context.Context, req *pricelist.GetProductPricesRequest) (*pricelist.GetProductPricesResponse, error) {
	ctx = log.ToContext(ctx, pls.logger)
	pls.logger.Info("get prices request received")
	sortBy := pricelist.SortingType_name[int32(req.SortingType)]
	sortDir := pricelist.SortingDirection_name[int32(req.SortingDirection)]

	var prices []types.ProductPriceExtended

	err := pls.runWithTimeout(ctx, pls.config.StoreTimeout, func(ctx context.Context) (err error) {
		pls.logger.Info("get prices from store")
		prices, err = pls.store.Get(ctx, types.SortingType(sortBy), types.SortDirectionType(sortDir), req.Limit, req.Offset)
		if err != nil {
			return errors.Wrap(err, "get data from store")
		}

		return nil
	})
	pls.logger.Info("prices fetched")
	if err != nil {
		return nil, err
	}

	response := pricelist.GetProductPricesResponse{Products: make([]*pricelist.ProductPrices, 0, len(prices))}
	for _, price := range prices {
		response.Products = append(response.Products, &pricelist.ProductPrices{
			ProductName:       price.ProductName,
			ProductPriceCents: price.ProductPriceCents,
			UpdateCount:       price.UpdatesCount,
			UpdateTime:        time.Unix(price.UpdateTimeNano/int64(time.Second), price.UpdateTimeNano%int64(time.Second)).Format(time.RFC3339Nano),
		})
	}
	return &response, nil
}

func (pls *PriceListServer) storePrices(ctx context.Context, prices []types.ProductPrice) error {
	return pls.runWithTimeout(ctx, pls.config.RestTimeout, func(ctx context.Context) (err error) {
		err = pls.store.Save(ctx, prices)
		if err != nil {
			// Error could be returned here. Depends on requirements
			pls.logger.WithError(err).Error("error to store requested data")
		}

		return nil
	})
}

func (pls *PriceListServer) fetchPrices(ctx context.Context, url string) (prices []types.ProductPrice, err error) {

	err = pls.runWithTimeout(ctx, pls.config.RestTimeout, func(ctx context.Context) (err error) {
		prices, err = pls.provider.Fetch(ctx, url)
		if err != nil {
			return errors.Wrap(err, "fetch data from url")
		}

		now := time.Now().UnixNano()
		for i := range prices {
			prices[i].UpdateTimeNano = now
		}
		return nil
	})

	return prices, err
}
