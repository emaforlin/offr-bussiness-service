package repository

import (
	"fmt"

	"github.com/emaforlin/bussiness-service/config"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.Config) (*gorm.DB, error) {
	var cfg = config.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ProvideRepo() fx.Option {
	return fx.Provide(InitDB)
}
