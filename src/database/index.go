package database

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/BananaFried525/home-restaurant-api/src/configs"
	"github.com/BananaFried525/home-restaurant-api/src/database/models"
)

var connection *gorm.DB

func Init() {
	_configs := configs.Load.Database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		_configs.User,
		_configs.Password,
		_configs.Host,
		_configs.Port,
		_configs.DbName,
	)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal(err)
	}

	// AUTO migrate for sync
	_db.AutoMigrate(
		&models.CustomerOrder{},
		&models.Customer{},
		&models.Food{},
		&models.Order{},
		&models.TableInfo{},
		&models.TableOrder{},
	)

	connection = _db
	log.Println("Database connected")
}

func Begin() *gorm.DB {
	log.Println("Database transaction begin")
	dbTxn := connection.Begin(&sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

	return dbTxn
}
