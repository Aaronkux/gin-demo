package global

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type CommonModel struct {
	ID        SnowflakeID           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"index"`
}
