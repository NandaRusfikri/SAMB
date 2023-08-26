package entities

import (
	"SAMB-BE/constant"
	"gorm.io/gorm"
	"time"
)

type TrxPenerimaanBarangDetail struct {
	TrxInDPK                  uint64                     `gorm:"column:TrxInDPK;primaryKey;AUTO_INCREMENT" json:"TrxInDPK"`
	CreatedAt                 time.Time                  `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt                 *time.Time                 `gorm:"updated_at" json:"-"`
	TrxInIDF                  uint                       `gorm:"column:TrxInIDF" json:"TrxInIDF"`
	TrxPenerimaanBarangHeader *TrxPenerimaanBarangHeader `gorm:"foreignKey:TrxInIDF" json:"TrxPenerimaanBarangHeader"`
	TrxInDProductIdf          uint                       `gorm:"column:TrxInDProductIdf" json:"TrxInDProductIdf"`
	MasterProduct             *MasterProduct             `gorm:"foreignKey:TrxInDProductIdf" json:"TrxInDProduct"`
	TrxInDQtyDus              int                        `gorm:"column:TrxInDQtyDus" json:"TrxInDQtyDus"`
	TrxInDQtyPcs              int                        `gorm:"column:TrxInDQtyPcs" json:"TrxInDQtyPcs"`
}

func (entity *TrxPenerimaanBarangDetail) TableName() string {
	return constant.TABLE_TRX_PENERIMAAN_BARANG_DETAIL
}

func (entity *TrxPenerimaanBarangDetail) BeforeUpdate(db *gorm.DB) error {
	time := time.Now()
	entity.UpdatedAt = &time
	return nil
}
