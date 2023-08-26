package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type EntityUnit struct {
	ID        uint64          `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	BlockName string          `gorm:"column:block_name;" json:"block_name,omitempty"`
	TowerName string          `gorm:"column:tower_name;" json:"tower_name,omitempty"`
	FloorName string          `gorm:"column:floor_name;" json:"floor_name,omitempty"`
	TypeName  string          `gorm:"column:type_name;" json:"type_name,omitempty"`
	//UnitCode string `gorm:"column:unit_code;" json:"unit_code,omitempty"`
	//UnitNo   string `gorm:"column:unit_no;" json:"unit_no,omitempty"`
	//Cluster  string `gorm:"column:cluster;" json:"cluster,omitempty"`
	//Address  string `gorm:"column:address;" json:"address"`
}

func (entity *EntityUnit) TableName() string {
	return constant.TABLE_UNIT_NAME
}

func (entity *EntityUnit) Validate() error {
	//if t.Code == "" {
	//	return fmt.Errorf("code required")
	//}

	return nil
}

func (entity *EntityUnit) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}
