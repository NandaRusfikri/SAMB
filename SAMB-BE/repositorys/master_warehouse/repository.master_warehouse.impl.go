package repositorys

import (
	"SAMB-BE/entities"
	"SAMB-BE/schemas"
	"gorm.io/gorm"
	"net/http"
)

type MasterWareHouseRepository interface {
	ListWareHouseRepository() ([]*entities.MasterWareHouse, schemas.SchemaDatabaseError)
}

type MasterWareHouseRepositoryImpl struct {
	db *gorm.DB
}

func InitMasterWareHouseRepository(dbCon *gorm.DB) MasterWareHouseRepository {
	return &MasterWareHouseRepositoryImpl{
		db: dbCon,
	}
}

func (r *MasterWareHouseRepositoryImpl) ListWareHouseRepository() ([]*entities.MasterWareHouse, schemas.SchemaDatabaseError) {

	var listWareHouse []*entities.MasterWareHouse
	Find := r.db.Find(&listWareHouse)
	if Find.Error != nil {
		return nil, schemas.SchemaDatabaseError{
			Code:  http.StatusBadRequest,
			Error: Find.Error,
		}
	}

	return listWareHouse, schemas.SchemaDatabaseError{}
}
