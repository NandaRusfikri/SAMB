package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type MasterSupplier struct {
	SupplierPK   uint64     `gorm:"column:SupplierPK;primaryKey;AUTO_INCREMENT" json:"SupplierPK"`
	CreatedAt    time.Time  `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt    *time.Time `gorm:"updated_at" json:"-"`
	SupplierName string     `gorm:"column:SupplierName" json:"SupplierName"`
}

func (entity *MasterSupplier) TableName() string {
	return constant.TABLE_MASTER_SUPPLIER
}

func (entity *MasterSupplier) BeforeUpdate(db *gorm.DB) error {
	time := time.Now()
	entity.UpdatedAt = &time
	return nil
}
