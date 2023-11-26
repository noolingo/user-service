package domain

type Config struct {
	AppName    string     `yaml:"appname" env-default:"test"`
	GrpcServer GrpcServer `yaml:"grpcserver"`
}

type GrpcServer struct {
	Host string `yaml:"host" env-default:"0.0.0.0"`
	Port string `yaml:"port" env-default:"9001"`
}
