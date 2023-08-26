package repositorys

import (
	"SAMB-BE/entities"
	"SAMB-BE/schemas"
)

type PenerimaanBarangRepository interface {
	CreatePenerimaanBarangHeaderRepository(input entities.TrxPenerimaanBarangHeader) (*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError)
	CreatePenerimaanBarangDetailRepository(input entities.TrxPenerimaanBarangDetail) (*entities.TrxPenerimaanBarangDetail, schemas.SchemaDatabaseError)
	ListBarangMasukHeaderRepository() ([]*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError)
	ListBarangMasukDetailRepository(TrxInPK uint) ([]*entities.TrxPenerimaanBarangDetail, schemas.SchemaDatabaseError)
	GetByIdBarangMasukHeaderRepository(TrxInPK uint) (*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError)
}
