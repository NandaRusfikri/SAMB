package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type TrxPengeluaranBarangDetail struct {
	TrxOutDPK uint64     `gorm:"column:TrxOutDPK;primaryKey;AUTO_INCREMENT" json:"TrxOutDPK"`
	CreatedAt time.Time  `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"-"`

	TrxOutIDF                  uint                        `gorm:"column:TrxOutIDF" json:"TrxOutIDF"`
	TrxPengeluaranBarangHeader *TrxPengeluaranBarangHeader `gorm:"foreignKey:TrxOutIDF" json:"TrxPengeluaranBarangHeader"`
	TrxOutDProductIdf          uint                        `gorm:"column:TrxOutDProductIdf" json:"TrxOutDProductIdf"`
	MasterProduct              *MasterProduct              `gorm:"foreignKey:TrxOutDProductIdf" json:"MasterProduct"`
	TrxOutDQtyDus              int                         `gorm:"column:TrxOutDQtyDus" json:"TrxOutDQtyDus"`
	TrxOutDQtyPcs              int                         `gorm:"column:TrxOutDQtyPcs" json:"TrxOutDQtyPcs"`
}

func (entity *TrxPengeluaranBarangDetail) TableName() string {
	return constant.TABLE_TRX_PENGELUARAN_BARANG_DETAIL
}

func (entity *TrxPengeluaranBarangDetail) BeforeUpdate(db *gorm.DB) error {
	time := time.Now()
	entity.UpdatedAt = &time
	return nil
}
