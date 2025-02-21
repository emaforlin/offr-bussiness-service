package repository

import (
	"fmt"

	"github.com/emaforlin/bussiness-service/config"
	"github.com/emaforlin/bussiness-service/entities"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

type MysqlDB struct {
	dsn    string
	logger logger.Interface
	db     *gorm.DB
}

func NewMySQLConnection(l *zap.Logger) GMDatabase {
	conf := config.GetConfig()

	logger := zapgorm2.New(l)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		conf.DB.User,
		conf.DB.Password,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.DBName,
	)

	return &MysqlDB{
		dsn:    dsn,
		logger: logger,
		db:     nil,
	}
}

func (m *MysqlDB) Connect() error {
	db, err := gorm.Open(mysql.Open(m.dsn), &gorm.Config{
		Logger: m.logger,
	})
	if err != nil {
		return err
	}
	m.db = db

	// Add your model to perform database migration
	if *config.GetConfig().DB.Migrate {
		db.AutoMigrate(entities.Business{}, entities.Employee{}, entities.Owner{})
	}

	return nil
}

func (m *MysqlDB) Cursor() *gorm.DB {
	return m.db
}
