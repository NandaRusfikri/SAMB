package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type TrxPengeluaranBarangHeader struct {
	TrxOutPK  uint64     `gorm:"column:TrxOutPK;primaryKey;AUTO_INCREMENT" json:"TrxOutPK"`
	CreatedAt time.Time  `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"-"`

	TrxOutNo        string           `gorm:"column:TrxOutNo" json:"TrxOutNo"`
	WhsIdf          uint             `gorm:"column:WhsIdf" json:"WhsIdf"`
	MasterWareHouse *MasterWareHouse `gorm:"foreignKey:WhsIdf;" json:"MasterWareHouse"`
	TrxOutDate      time.Time        `gorm:"column:TrxOutDate" json:"TrxOutDate"`
	TrxOutSuppIdf   uint             `gorm:"column:TrxOutSuppIdf" json:"TrxOutSuppIdf"`
	MasterSupplier  *MasterSupplier  `gorm:"foreignKey:TrxOutSuppIdf" json:"MasterSupplier"`
	TrxOutNotes     string           `gorm:"column:TrxOutNotes" json:"TrxOutNotes"`
}

func (entity *TrxPengeluaranBarangHeader) TableName() string {
	return constant.TABLE_TRX_PENGELUARAN_BARANG_HEADER
}

func (entity *TrxPengeluaranBarangHeader) BeforeUpdate(db *gorm.DB) error {
	time := time.Now()
	entity.UpdatedAt = &time
	return nil
}
