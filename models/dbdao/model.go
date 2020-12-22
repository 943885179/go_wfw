package db

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
	Updated   string         `gorm:"autoUpdateTime:milli" json:"updated"` // 使用时间戳毫秒数填充更新时间
	Created   string         `gorm:"autoCreateTime" json:"created"`       // 使用时间戳秒数填充创建时间
}
