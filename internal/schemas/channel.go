package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Channel struct {
	ID          uuid.UUID
	ChannelName string
	Password    string
	OwnerSlug   string
	ManagerName string
	UpdatedAt   time.Time
	CreatedAt   time.Time
	TTL         time.Time
}

func (ch *Channel) BeforeCreate(tx *gorm.DB) (err error) {
	ch.ID = uuid.New()
	ch.TTL = time.Now().Add(time.Minute)
	return
}
