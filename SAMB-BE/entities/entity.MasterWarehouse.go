package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type MasterWareHouse struct {
	WhsPK     uint64     `gorm:"column:WhsPK;primaryKey;AUTO_INCREMENT" json:"WhsPK"`
	CreatedAt time.Time  `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"-"`
	WhsName   string     `gorm:"column:WhsName" json:"WhsName"`
}

func (entity *MasterWareHouse) TableName() string {
	return constant.TABLE_MASTER_WAREHOUSE
}

func (entity *MasterWareHouse) BeforeUpdate(db *gorm.DB) error {
	time := time.Now()
	entity.UpdatedAt = &time
	return nil
}
