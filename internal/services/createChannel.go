package services

import (
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
	}

	res := db.Db.Create(&channel)
	return res.Error
}
