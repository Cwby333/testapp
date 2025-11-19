package internal

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Host        string `yaml:"host"`
	Port        uint16 `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	DB          string `yaml:"database"`
	MaxConns    int `yaml:"max-conns"`
	MinConns    int `yaml:"min-conns"`
	MaxIdleConn time.Duration `yaml:"max-idle-conn"`
	MaxLifetimeConn time.Duration `yaml:"max-lifetime-conn"`
}

func Load() Config {
	cfg := Config{}

	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}

	err = cleanenv.ReadConfig(os.Getenv("TESTAPP_CONFIG_PATH"), &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}