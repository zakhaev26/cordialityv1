package services

import (
	"time"

	"github.com/fuckthinkpad/internal/db"
	"github.com/fuckthinkpad/internal/schemas"
)

func GetChannelService(channelName string) (schemas.Channel, error) {

	var ch schemas.Channel

	if res := db.Db.Where("channel_name = ?",channelName).Where("ttl > ?",time.Now()).First(&ch); res.Error != nil {
		return schemas.Channel{}, res.Error
	}
	return ch, nil
}
