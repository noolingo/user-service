package domain

type Config struct {
	AppName    string     `yaml:"appname" env-default:"test"`
	GrpcServer GrpcServer `yaml:"grpcserver"`
	Mysql      Mysql      `yaml:"mysql" env-prefix:"USER_SERVICE_"`
}

type GrpcServer struct {
	Host string `yaml:"host" env-default:"0.0.0.0"`
	Port string `yaml:"port" env-default:"9001"`
}

type Mysql struct {
	DSN string `yaml:"dsn" env:"MYSQL_DSN"`
}
