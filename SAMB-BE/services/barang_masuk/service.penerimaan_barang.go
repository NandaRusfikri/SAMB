package services

import (
	"SAMB-BE/entities"
	"SAMB-BE/schemas"
)

type ServicePenerimaanBarang interface {
	CreatePenerimaanBarangHeaderService(input schemas.PenerimaanBarangHeaderRequest) (*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError)

	CreatePenerimaanBarangDetailService(input schemas.PenerimaanBarangDetailRequest) (*entities.TrxPenerimaanBarangDetail, schemas.SchemaDatabaseError)

	ListBarangMasukHeaderService() ([]*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError)
	ListBarangMasukDetailService(TrxInPK uint) ([]*entities.TrxPenerimaanBarangDetail, schemas.SchemaDatabaseError)

	GetByIdBarangMasukHeaderService(TrxInPK uint) (*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError)
}
