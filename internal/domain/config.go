package domain

import "time"

type Config struct {
	AppName    string     `yaml:"appname" env-default:"test"`
	GrpcServer GrpcServer `yaml:"grpcserver"`
	Mysql      Mysql      `yaml:"mysql" env-prefix:"USER_SERVICE_"`
	Auth       AppAuth    `yaml:"auth" env-prefix:"AUTH_"`
}

type GrpcServer struct {
	Host string `yaml:"host" env-default:"0.0.0.0"`
	Port string `yaml:"port" env-default:"9001"`
}

type Mysql struct {
	DSN             string        `yaml:"dsn" env:"MYSQL_DSN"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" env-default:"5m"`
	MaxOpenConns    int           `yaml:"max_open_conns" env-default:"10"`
	MaxIdleConns    int           `yaml:"max_idle_conns" env-default:"10"`
}

type AppAuth struct {
	AccessSecretKey string        `yaml:"access-secret-key" env:"ACCESS_SECRET_KEY"`
	AccessKeyPrefix string        `yaml:"access-key-prefix" env-default:"access-key"`
	AccessKeyTtl    time.Duration `yaml:"access-key-ttl" env:"ACCESS_KEY_TTL" env-default:"15m"`

	RefreshSecretKey string        `yaml:"refresh-secret-key" env:"REFRESH_SECRET_KEY"`
	RefreshKeyPrefix string        `yaml:"refresh-key-prefix" env-default:"refresh-key"`
	RefreshKeyTtl    time.Duration `yaml:"refresh-key-ttl" env:"REFRESH_KEY_TTL" env-default:"120h"`
}
