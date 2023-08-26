package entities

import (
	"SAMB-BE/constant"
	"SAMB-BE/util/audit"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserActivityLog struct {
	audit.CreateAudit
	//FeatureId int64             `gorm:"column:feature_id;" json:"feature_id"`
	//Feature   *mFeature.Feature `gorm:"foreignKey:feature_id" json:"feature"`
	UserId   int64       `gorm:"column:user_id;" json:"user_id,omitempty" `
	User     *EntityUser `gorm:"foreignKey:user_id" json:"user,omitempty"`
	Activity string      `gorm:"column:activity" json:"activity"`
	DataId   int64       `gorm:"column:data_id" json:"data_id"`
	DataRaw  string      `gorm:"column:data_raw" json:"data_raw"`
}

func (t *UserActivityLog) TableName() string {
	return constant.TABLE_USER_ACTIVITY_LOG_NAME
}

func (t *UserActivityLog) Validate() error {

	if t.UserId == 0 {
		return fmt.Errorf("user_id required")
	}

	return nil
}

func (t *UserActivityLog) InitAudit(operation string, user int64) {
	timeNow := time.Now()
	switch operation {
	case constant.OPERATION_SQL_INSERT:
		t.CreatedUser = user
		t.CreatedAt = timeNow

	}
}

func (UserActivityLog) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&UserActivityLog{})
	if err != nil {
		return err
	}

	return nil
}
