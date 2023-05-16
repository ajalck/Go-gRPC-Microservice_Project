package db

import (
	"log"

	"github.com/ajalck/Go-gRPC-Microservice_Project/order_management/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func InitDB(url string) *UserDB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	return &UserDB{DB: db}
}
func SyncDB(db *gorm.DB) (err error) {
	if err = (db.AutoMigrate(&models.Order{})); err != nil {
		return err
	}
	return nil
}
