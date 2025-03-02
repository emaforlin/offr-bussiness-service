package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Config struct {
	App App
	DB  DB
}

type App struct {
	Port    uint16
	DevMode bool
}

type DB struct {
	DBName   string
	Port     uint16
	Host     string
	User     string
	Password string
	Migrate  *bool
}

var config *Config

func Init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".config/")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error parsing default config file")
	}

	config = &Config{
		App: App{
			viper.GetUint16("service.port"),
			viper.GetBool("service.devmode"),
		},
		DB: DB{
			DBName:   viper.GetString("database.dbname"),
			Host:     viper.GetString("database.host"),
			Port:     viper.GetUint16("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			Migrate:  flag.Bool("migrate", false, "set to true to migrate database schema"),
		},
	}
	flag.Parse()
}

func GetConfig() *Config {
	return config
}

func ProvideConfig() fx.Option {
	Init()
	return fx.Provide(GetConfig)
}
