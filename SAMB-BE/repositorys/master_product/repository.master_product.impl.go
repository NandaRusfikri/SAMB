package repositorys

import (
	"SAMB-BE/entities"
	"SAMB-BE/schemas"
	"gorm.io/gorm"
	"net/http"
)

type MasterProductRepository interface {
	ListProductRepository() ([]*entities.MasterProduct, schemas.SchemaDatabaseError)
}

type MasterProductRepositoryImpl struct {
	db *gorm.DB
}

func InitMasterProductRepository(dbCon *gorm.DB) MasterProductRepository {
	return &MasterProductRepositoryImpl{
		db: dbCon,
	}
}

func (r *MasterProductRepositoryImpl) ListProductRepository() ([]*entities.MasterProduct, schemas.SchemaDatabaseError) {

	var listProduct []*entities.MasterProduct
	Find := r.db.Find(&listProduct)
	if Find.Error != nil {
		return nil, schemas.SchemaDatabaseError{
			Code:  http.StatusBadRequest,
			Error: Find.Error,
		}
	}

	return listProduct, schemas.SchemaDatabaseError{}
}
