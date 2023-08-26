package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type TasklistLogImport struct {
	ID        uint64     `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"-"`
	Period    string     `gorm:"column:period;" json:"period"`

	UserId *int64      `gorm:"column:user_id;" json:"user_id,omitempty"`
	User   *EntityUser `gorm:"foreignKey:user_id;references:id" json:"user,omitempty"`
	Status string      `gorm:"column:status;" json:"status"`
	File   string      `gorm:"column:file;" json:"file"`
}

func (t *TasklistLogImport) TableName() string {
	return constant.TABLE_TASKLIST_LOG_IMPORT_NAME
}

func (t *TasklistLogImport) Validate() error {
	//if t.Code == "" {
	//	return fmt.Errorf("code required")
	//}

	return nil
}

func (entity *TasklistLogImport) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}
