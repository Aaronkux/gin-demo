package global

import (
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	ID        SnowflakeID    `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
