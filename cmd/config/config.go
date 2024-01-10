package config

type Config struct {
	HttpPort int `mapstructure:"httpPort"`
	GrpcPort int `mapstructure:"grpcPort"`
}
