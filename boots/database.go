package boots

import (
	"fmt"
	"log"

	"github.com/BananaFried525/home-restaurant-api/internal/adapters/repository/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase() *gorm.DB {
	_configs := Config.Database
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

	log.Println("Database connected")

	return _db
}
