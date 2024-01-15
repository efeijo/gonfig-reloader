package config

type Config struct {
	GrpcServePort  int `mapstructure:"grpcServePort"`
	HttpServerPort int `mapstructure:"httpServerPort"`
}
