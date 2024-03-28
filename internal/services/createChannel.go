package services

import (
	"time"

	"github.com/fuckthinkpad/internal/db"
	"github.com/fuckthinkpad/internal/schemas"
)

func CreateChannelService(reqBody struct {
	ChannelName string `json:"channelName,omitempty"`
	Password    string `json:"password,omitempty"`
	OwnerSlug   string `json:"ownerSlug,omitempty"`
}, managerName string) error {

	channel := schemas.Channel{
		ChannelName: reqBody.ChannelName,
		Password:    reqBody.Password,
		ManagerName: managerName,
		OwnerSlug:   reqBody.OwnerSlug,
		TTL: time.Now(),
	}

	res := db.Db.Create(&channel)
	return res.Error
}
