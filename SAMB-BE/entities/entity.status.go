package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type EntityStatus struct {
	ID        uint64     `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`

	Name string `gorm:"column:name;" json:"name" example:"Tidak Jadi Beli"`
}

func (t *EntityStatus) TableName() string {
	return constant.TABLE_STATUS_NAME
}

func (entity *EntityStatus) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}
