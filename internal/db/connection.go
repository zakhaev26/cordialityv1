package db

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/fuckthinkpad/internal/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {

	var err error
	// https://github.com/go-gorm/postgres
	Db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=soubhik password=soubhik dbname=distributed_chat_app host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Error("Failed to connect to Database", "err", err)
		os.Exit(1)
	}
	log.Info("Connected to PostgreSQL")
	if err := Db.AutoMigrate(&schemas.Channel{}); err != nil {
		log.Error("Schema Migration Failed", "err", err)
		return
	}
	log.Info("Migrations Successful")
}
