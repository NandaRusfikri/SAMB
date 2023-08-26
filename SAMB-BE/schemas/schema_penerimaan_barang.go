package schemas

import (
	"SAMB-BE/entities"
	"time"
)

type PenerimaanBarangHeaderRequest struct {
	TrxInNo      string ` json:"TrxInNo" example:"TRX-001"`
	WhsIdf       uint   ` json:"WhsIdf" example:"1"`
	TrxInSuppIdf uint   ` json:"TrxInSuppIdf" example:"1"`
	TrxInNotes   string ` json:"TrxInNotes" example:"catatan"`
}
type PenerimaanBarangDetailRequest struct {
	TrxInIDF         uint ` json:"TrxInIDF"`
	TrxInDProductIdf uint `json:"TrxInDProductIdf"`
	TrxInDQtyDus     int  ` json:"TrxInDQtyDus"`
	TrxInDQtyPcs     int  ` json:"TrxInDQtyPcs"`
}
type DetailPenerimaanBarangResponse struct {
	TrxInPK                   uint64                                `gorm:"primaryKey;column:TrxInPK" json:"TrxInPK"`
	CreatedAt                 time.Time                             `gorm:"created_at;default:now()" json:"created_at,omitempty"`
	UpdatedAt                 *time.Time                            `gorm:"updated_at" json:"-"`
	TrxInNo                   string                                `gorm:"column:TrxInNo" json:"TrxInNo"`
	WhsIdf                    uint                                  `gorm:"column:WhsIdf" json:"WhsIdf"`
	MasterWareHouse           *entities.MasterWareHouse             `gorm:"foreignKey:WhsIdf"`
	TrxInDate                 time.Time                             `gorm:"column:TrxInDate" json:"TrxInDate"`
	TrxInSuppIdf              uint                                  `gorm:"column:TrxInSuppIdf" json:"TrxInSuppIdf"`
	MasterSupplier            *entities.MasterSupplier              `gorm:"foreignKey:TrxInSuppIdf"`
	TrxInNotes                string                                `gorm:"column:TrxInNotes" json:"TrxInNotes"`
	TrxPenerimaanBarangDetail []*entities.TrxPenerimaanBarangDetail ` json:"TrxPenerimaanBarangDetail"`
}

type SchemaListPenerimaanBarang struct {
	SearchText  string `json:"search_text" form:"search_text" example:"Search example name,email or code"`
	OrderField  string `json:"order_field" form:"order_field" example:"id|desc"`
	FilterPage  int    `json:"filter_page" form:"filter_page" example:"1"`
	FilterLimit int    `json:"filter_limit" form:"filter_limit" example:"10"`

	BlockName string ` json:"block_name"`
	TowerName string ` json:"tower_name"`
	FloorName string ` json:"floor_name"`
	TypeName  string ` json:"type_name"`
}
