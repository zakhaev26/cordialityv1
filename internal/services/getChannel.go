package services

import (
	"github.com/fuckthinkpad/internal/db"
	"github.com/fuckthinkpad/internal/schemas"
)

func GetChannelService(reqBody struct {
	ChannelName string `json:"channelName,omitempty"`
	Password    string `json:"password,omitempty"`
}) (schemas.Channel, error) {

	var ch schemas.Channel

	if res := db.Db.Where(&schemas.Channel{
		ChannelName: reqBody.ChannelName,
	}).First(&ch); res.Error != nil {
		return schemas.Channel{}, res.Error
	}

	return ch, nil
}
