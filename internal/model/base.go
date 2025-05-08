package model

import (
	"time"

	"39.com/pkg/database"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID     uint      `gorm:"primarykey"`
	Status uint8     `gorm:"default:1"`
	Ctime  time.Time `gorm:"autoCreateTime"`
	Uptime time.Time `gorm:"autoUpdateTime"`
	Db     *gorm.DB
}

func (m *BaseModel) GetDB() *gorm.DB {
	if m.Db == nil {
		m.Db = database.GetDb()
	}
	return m.Db
}
