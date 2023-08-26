package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type EntityUser struct {
	ID        uint64     `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time  `gorm:"created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"-"`
	DeletedAt *time.Time `gorm:"deleted_at" json:"-"`
	Name      string     `gorm:"name" json:"name"`
	//Email      string     `gorm:"email" json:"email"`
	//Phone      string     `gorm:"phone" json:"phone,omitempty"`
	Username   string `gorm:"username" json:"username,omitempty"`
	Password   string `gorm:"password" json:"-,omitempty"`
	IsActive   *bool  `gorm:"is_active" json:"is_active,omitempty"`
	AvatarPath string `gorm:"avatar_path" json:"avatar_path,omitempty"`
	//Token      string     `gorm:"column:token;" json:"token,omitempty"`
}

func (entity *EntityUser) TableName() string {
	return constant.TABLE_USERS
}
func (entity *EntityUser) BeforeCreate(db *gorm.DB) error {
	//model.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityUser) BeforeUpdate(db *gorm.DB) error {
	time := time.Now().Local()
	entity.UpdatedAt = &time
	return nil
}
