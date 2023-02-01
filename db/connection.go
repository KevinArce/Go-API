package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=root password=mysecretpassword dbname=gorm port=5432"
var DB *gorm.DB
var err error

func Connect() (*gorm.DB, error) {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		println("Connected to database")
	}
	return DB, nil
}