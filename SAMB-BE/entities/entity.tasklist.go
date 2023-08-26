package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type EntityTasklist struct {
	ID            uint64      `gorm:"column:id;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt     time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     *time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     *time.Time  `gorm:"column:deleted_at" json:"-"`
	Period        string      `gorm:"column:period;" json:"period"`
	UnitId        uint64      `gorm:"column:unit_id;" json:"unit_id,omitempty"`
	Unit          *EntityUnit `gorm:"foreignKey:unit_id;references:id" json:"unit,omitempty"`
	UserId        *uint64     `gorm:"column:user_id;" json:"user_id,omitempty"`
	User          *EntityUser `gorm:"foreignKey:user_id;references:id" json:"user,omitempty"`
	ScheduleVisit *time.Time  `gorm:"column:schedule_visit" json:"schedule_visit"`
	Shift         string      `gorm:"column:shift" json:"shift"`
	//TasklistMeter []*EntityTasklistMeter `gorm:"foreignKey:tasklist_id" json:"tasklist_meter,omitempty"`

	ProblemId       *uint64        `gorm:"column:problem_id;" json:"problem_id,omitempty"`
	Problem         *EntityProblem `gorm:"foreignKey:problem_id" json:"problem,omitempty"`
	TypeId          int            `gorm:"column:type_id;" json:"type_id,omitempty"`
	PrevRead        int64          `gorm:"column:prev_read;" json:"prev_read"`
	CurrRead        int64          `gorm:"column:curr_read;" json:"curr_read"`
	Consumption     int64          `gorm:"-" json:"consumption"`
	CurrConsumption int64          `gorm:"-" json:"curr_consumption"`
	CurrCorrection  int64          `gorm:"column:curr_correction;" json:"curr_correction"`
	StatusId        uint64         `gorm:"column:status_id;default:1" json:"status_id"`
	Status          *EntityStatus  `gorm:"foreignKey:status_id" json:"status"`
	PathPhotoNumber string         `gorm:"column:path_photo_number;" json:"path_photo_number"`
	PathPhotoUnit   string         `gorm:"column:path_photo_unit;" json:"path_photo_unit"`
	URLPhotoNumber  string         `gorm:"-" sql:"-" json:"url_photo_number"`
	URLPhotoUnit    string         `gorm:"-" sql:"-" json:"url_photo_unit"`
	//Threshold         string                `gorm:"-" sql:"-" json:"threshold"`
	VisitDate *time.Time `gorm:"column:visit_date;" json:"visit_date"`
}

func (entity *EntityTasklist) TableName() string {
	return constant.TABLE_TASKLIST_NAME
}

func (entity *EntityTasklist) Validate() error {

	return nil
}

func (entity *EntityTasklist) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}
