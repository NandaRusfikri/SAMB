package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type MasterCustomer struct {
	CustomerPK   uint64     `gorm:"column:CustomerPK;primaryKey;AUTO_INCREMENT" json:"CustomerPK"`
	CreatedAt    time.Time  `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt    *time.Time `gorm:"updated_at" json:"-"`
	CustomerName string     `gorm:"column:CustomerName" json:"CustomerName"`
}

func (entity *MasterCustomer) TableName() string {
	return constant.TABLE_MASTER_CUSTOMER
}

func (entity *MasterCustomer) BeforeUpdate(db *gorm.DB) error {
	time := time.Now()
	entity.UpdatedAt = &time
	return nil
}
