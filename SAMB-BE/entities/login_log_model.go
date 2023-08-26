package entities

import (
	"SAMB-BE/constant"
	"SAMB-BE/util/audit"
	"gorm.io/gorm"
	"time"
)

type UserLog struct {
	audit.FullAudit
	DeviceId string      `gorm:"column:device_id" json:"device_id"`
	UserId   int64       `gorm:"column:user_id;" json:"user_id,omitempty" `
	User     *EntityUser `gorm:"foreignKey:user_id" json:"user,omitempty"`
	//MerchantId *int64              `gorm:"column:merchant_id;" json:"merchant_id,omitempty" `
	//Merchant   *mMerchant.Merchant `gorm:"foreignKey:merchant_id" json:"merchant,omitempty"`
}

func (t *UserLog) TableName() string {
	return constant.TABLE_USER_LOG_NAME
}

func (t *UserLog) Validate() error {

	//if t.ProductId == 0 {
	//	return fmt.Errorf("product_id required")
	//}

	return nil
}

func (t *UserLog) InitAudit(operation string, user int64) {
	timeNow := time.Now()
	switch operation {
	case constant.OPERATION_SQL_INSERT:
		t.CreatedUser = user
		t.CreatedAt = &timeNow
		t.UpdatedUser = user
		t.UpdatedAt = &timeNow
	case constant.OPERATION_SQL_UPDATE:
		t.UpdatedUser = user
		t.UpdatedAt = &timeNow
	case constant.OPERATION_SQL_DELETE:
		t.DeletedUser = user
		t.DeletedAt = &gorm.DeletedAt{Time: timeNow, Valid: true}
	}
}

func (UserLog) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&UserLog{})
	if err != nil {
		return err
	}

	return nil
}
