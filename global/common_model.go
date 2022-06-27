package global

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type CommonModel struct {
	ID        SnowflakeID           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time             `json:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
	DeletedAt soft_delete.DeletedAt `json:"-"`
}
