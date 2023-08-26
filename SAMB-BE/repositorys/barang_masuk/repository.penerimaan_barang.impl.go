package repositorys

import (
	"SAMB-BE/entities"
	"SAMB-BE/schemas"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

type PenerimaanBarangRepositoryImpl struct {
	db *gorm.DB
}

func InitPenerimaanBarangRepository(dbCon *gorm.DB) PenerimaanBarangRepository {
	return &PenerimaanBarangRepositoryImpl{
		db: dbCon,
	}
}

func (r *PenerimaanBarangRepositoryImpl) CreatePenerimaanBarangHeaderRepository(input entities.TrxPenerimaanBarangHeader) (*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError) {

	Create := r.db.Create(&input)
	if Create.Error != nil {
		return nil, schemas.SchemaDatabaseError{
			Code:  http.StatusBadRequest,
			Error: Create.Error,
		}
	}

	return &input, schemas.SchemaDatabaseError{}
}
func (r *PenerimaanBarangRepositoryImpl) CreatePenerimaanBarangDetailRepository(input entities.TrxPenerimaanBarangDetail) (*entities.TrxPenerimaanBarangDetail, schemas.SchemaDatabaseError) {

	Create := r.db.Create(&input)
	if Create.Error != nil {
		return nil, schemas.SchemaDatabaseError{
			Code:  http.StatusBadRequest,
			Error: Create.Error,
		}
	}

	return &input, schemas.SchemaDatabaseError{}
}

func (r *PenerimaanBarangRepositoryImpl) ListBarangMasukHeaderRepository() ([]*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError) {

	var data []*entities.TrxPenerimaanBarangHeader
	Find := r.db.Preload(clause.Associations).Find(&data)
	if Find.Error != nil {
		return nil, schemas.SchemaDatabaseError{
			Code:  http.StatusBadRequest,
			Error: Find.Error,
		}
	}

	return data, schemas.SchemaDatabaseError{}
}

func (r *PenerimaanBarangRepositoryImpl) GetByIdBarangMasukHeaderRepository(TrxInPK uint) (*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError) {

	var data entities.TrxPenerimaanBarangHeader
	Find := r.db.Preload(clause.Associations).Where(`"TrxInPK" = ?`, TrxInPK).First(&data)
	if Find.Error != nil {
		return nil, schemas.SchemaDatabaseError{
			Code:  http.StatusBadRequest,
			Error: Find.Error,
		}
	}

	return &data, schemas.SchemaDatabaseError{}
}

func (r *PenerimaanBarangRepositoryImpl) ListBarangMasukDetailRepository(TrxInPK uint) ([]*entities.TrxPenerimaanBarangDetail, schemas.SchemaDatabaseError) {

	var data []*entities.TrxPenerimaanBarangDetail
	Find := r.db.Preload(clause.Associations).Where(`"TrxInIDF" = ?`, TrxInPK).Find(&data)
	if Find.Error != nil {
		return nil, schemas.SchemaDatabaseError{
			Code:  http.StatusBadRequest,
			Error: Find.Error,
		}
	}

	return data, schemas.SchemaDatabaseError{}
}
