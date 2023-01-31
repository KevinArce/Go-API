package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=root password=mysecretpassword dbname=gorm port=5432"
var db *gorm.DB
var err error

func Connect() (*gorm.DB, error) {
	db, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}