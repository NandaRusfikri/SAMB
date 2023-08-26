package repositorys

import (
	"SAMB-BE/entities"
	"SAMB-BE/schemas"
	"gorm.io/gorm"
	"net/http"
)

type MasterSupplierRepository interface {
	ListSupplierRepository() ([]*entities.MasterSupplier, schemas.SchemaDatabaseError)
}

type MasterSupplierRepositoryImpl struct {
	db *gorm.DB
}

func InitMasterSupplierRepository(dbCon *gorm.DB) MasterSupplierRepository {
	return &MasterSupplierRepositoryImpl{
		db: dbCon,
	}
}

func (r *MasterSupplierRepositoryImpl) ListSupplierRepository() ([]*entities.MasterSupplier, schemas.SchemaDatabaseError) {

	var listSupplier []*entities.MasterSupplier
	Find := r.db.Find(&listSupplier)
	if Find.Error != nil {
		return nil, schemas.SchemaDatabaseError{
			Code:  http.StatusBadRequest,
			Error: Find.Error,
		}
	}

	return listSupplier, schemas.SchemaDatabaseError{}
}
