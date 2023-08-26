package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type MasterProduct struct {
	ProductPK   uint64     `gorm:"column:ProductPK;primaryKey;AUTO_INCREMENT" json:"ProductPK"`
	CreatedAt   time.Time  `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"updated_at" json:"-"`
	ProductName string     `gorm:"column:ProductName" json:"ProductName"`
}

func (entity *MasterProduct) TableName() string {
	return constant.TABLE_MASTER_PRODUCT
}

func (entity *MasterProduct) BeforeUpdate(db *gorm.DB) error {
	time := time.Now()
	entity.UpdatedAt = &time
	return nil
}
