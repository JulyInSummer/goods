package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App   App        `envconfig:"APP"`
	DB    DB         `envconfig:"DB"`
	Redis Redis      `envconfig:"REDIS"`
	Ch    Clickhouse `envconfig:"CLICKHOUSE"`
}

type App struct {
	Env     string        `envconfig:"ENV"`
	Host    string        `envconfig:"HOST"`
	Port    string        `envconfig:"PORT"`
	Timeout time.Duration `envconfig:"TIMEOUT"`
}

type DB struct {
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	Name     string `envconfig:"NAME"`
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
	SSL      string `envconfig:"SSL"`
}

type Redis struct {
	Host string `envconfig:"HOST"`
	Port string `envconfig:"PORT"`
	DB   int    `envconfig:"DB"`
}

type Clickhouse struct {
	DB       string `envconfig:"DB"`
	Username string `envconfig:"USERNAME"`
	Password string `envconfig:"PASSWORD"`
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
}

func NewConfig() *Config {
	if err := godotenv.Load("./config/.env"); err != nil {
		panic(err)
	}

	var cfg Config
	envconfig.MustProcess("", &cfg)

	return &cfg
}
