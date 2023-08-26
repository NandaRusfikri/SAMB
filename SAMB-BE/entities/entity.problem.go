package entities

import (
	"SAMB-BE/constant"
	"time"

	"gorm.io/gorm"
)

type EntityProblem struct {
	ID        uint64          `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time       `gorm:"column:created_at" json:"-"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"-"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at" json:"-" `
	IsActive  *bool           `gorm:"column:is_active;default:true" json:"is_active"`
	Name      string          `gorm:"column:name" json:"name"`
}

func (entity *EntityProblem) TableName() string {
	return constant.TABLE_PROBLEM
}
func (entity *EntityProblem) BeforeCreate(db *gorm.DB) error {
	//model.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityProblem) BeforeUpdate(db *gorm.DB) error {
	time := time.Now().Local()
	entity.UpdatedAt = &time
	return nil
}
