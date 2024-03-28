package db

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/fuckthinkpad/internal/schemas"
	"github.com/fuckthinkpad/internal/ws"
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

	//re-generate all the channels
	channels := FindAllChannels()

	for _, v := range channels {
		// if time.Until(v.TTL) <  {

		// }
		ws.MasterManager.SetManager(v.ManagerName, ws.NewManager(v.ManagerName))
		log.Info("Reincarted", "Channel", v.ChannelName)
	}
}

func FindAllChannels() []schemas.Channel {
	var channels []schemas.Channel

	if res := Db.Where("ttl > ?", time.Now()).Find(&channels); res.Error != nil {
		log.Warn("Reincarnation Failed!")
		os.Exit(1)
	}

	return channels
}
