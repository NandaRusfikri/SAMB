package services

import (
	"SAMB-BE/entities"
	repositorys "SAMB-BE/repositorys/master_warehouse"
	"SAMB-BE/schemas"
)

type ServiceMasterWareHouse interface {
	ListMasterWareHouseService() ([]*entities.MasterWareHouse, schemas.SchemaDatabaseError)
}

type ServiceMasterWareHouseImpl struct {
	MasterWareHouseRepository repositorys.MasterWareHouseRepository
}

func InitServiceMasterWareHouse(unitRepository repositorys.MasterWareHouseRepository) ServiceMasterWareHouse {
	return &ServiceMasterWareHouseImpl{
		MasterWareHouseRepository: unitRepository,
	}
}

func (s *ServiceMasterWareHouseImpl) ListMasterWareHouseService() ([]*entities.MasterWareHouse, schemas.SchemaDatabaseError) {

	res, err := s.MasterWareHouseRepository.ListWareHouseRepository()

	return res, err
}
