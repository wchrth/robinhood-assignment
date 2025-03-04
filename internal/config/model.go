package config

import "time"

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	SecretKey             string        `mapstructure:"secret_key"`
	Issuer                string        `mapstructure:"issuer"`
	AccessExpirationTime  time.Duration `mapstructure:"access_expiration_time"`
	RefreshExpirationTime time.Duration `mapstructure:"refresh_expiration_time"`
}
