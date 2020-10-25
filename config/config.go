package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Config struct {
	Mongo        MongoConfig
	GRPC         GRPCConfig
	ServerConfig ServerConfig
	HTTP         HTTPTestServer
}

type MongoConfig struct {
	URL            string
	DB             string
	Collection     string
	ConnectTimeout time.Duration
}

type GRPCConfig struct {
	Port    int
	Options []grpc.ServerOption
}

type ServerConfig struct {
	StoreTimeout time.Duration
	RestTimeout  time.Duration
}

type HTTPTestServer struct {
	Port         int
	FilePrefix   string
	RequestCount int
}

func ParseFromEnvs(envs []string) (config Config, err error) {
	opts := make([]grpc.ServerOption, 0)
	// TODO: parse isTLS from flags
	isTLS := false
	if isTLS {
		certFileName := ""
		keyFileName := ""
		creds, err := credentials.NewServerTLSFromFile(certFileName, keyFileName)
		if err != nil {
			return config, errors.Wrapf(err, "failed to generate credentials. CertFile:%v KeyFile: %v", certFileName, keyFileName)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	err = envconfig.Process("app", &config)
	return config, err

	//config = Config{
	//	Mongo: MongoConfig{
	//		URL:            "mongodb://localhost:27017",
	//		DB:             "prices",
	//		Collection:     "prices",
	//		ConnectTimeout: 5 * time.Second,
	//	},
	//	GRPC: GRPCConfig{
	//		HostPort: "0.0.0.0:3000",
	//		Options:  opts,
	//	},
	//	ServerConfig: ServerConfig{
	//		StoreTimeout: time.Second,
	//		RestTimeout:  5 * time.Second,
	//	},
	//}
	//
	//return config, nil
}
