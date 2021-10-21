package storage

import (
	"fmt"
	"log"

	config "web-server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable",
			config.DBUser, config.DBPassword, config.DBName, config.DBHost, config.DBPort),
		PreferSimpleProtocol: true,
	}))

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
