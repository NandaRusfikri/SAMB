package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type TrxPenerimaanBarangHeader struct {
	TrxInPK         uint64           `gorm:"primaryKey;column:TrxInPK" json:"TrxInPK"`
	CreatedAt       time.Time        `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt       *time.Time       `gorm:"updated_at" json:"-"`
	TrxInNo         string           `gorm:"column:TrxInNo" json:"TrxInNo"`
	WhsIdf          uint             `gorm:"column:WhsIdf" json:"WhsIdf"`
	MasterWareHouse *MasterWareHouse `gorm:"foreignKey:WhsIdf"`
	TrxInDate       time.Time        `gorm:"column:TrxInDate" json:"TrxInDate"`
	TrxInSuppIdf    uint             `gorm:"column:TrxInSuppIdf" json:"TrxInSuppIdf"`
	MasterSupplier  *MasterSupplier  `gorm:"foreignKey:TrxInSuppIdf"`
	TrxInNotes      string           `gorm:"column:TrxInNotes" json:"TrxInNotes"`
	//TrxPenerimaanBarangDetail []*TrxPenerimaanBarangDetail ` json:"TrxPenerimaanBarangDetail"`
}

func (entity *TrxPenerimaanBarangHeader) TableName() string {
	return constant.TABLE_TRX_PENERIMAAN_BARANG_HEADER
}

func (entity *TrxPenerimaanBarangHeader) BeforeUpdate(db *gorm.DB) error {
	time := time.Now()
	entity.UpdatedAt = &time
	return nil
}
