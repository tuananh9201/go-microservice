package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func InitDB(c *DBConfig) (*gorm.DB, error) {
	// var db *gorm.DB
	dsn := "host=" + c.Host + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName + " port=" + c.Port + " sslmode=" + c.SSLMode
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	fmt.Println("Migrated database")
	return db, nil
}
