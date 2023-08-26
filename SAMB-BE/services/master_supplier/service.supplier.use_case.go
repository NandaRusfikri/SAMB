package services

import (
	"SAMB-BE/entities"
	repositorys "SAMB-BE/repositorys/master_supplier"
	"SAMB-BE/schemas"
)

type ServiceMasterSupplier interface {
	ListMasterSupplierService() ([]*entities.MasterSupplier, schemas.SchemaDatabaseError)
}

type ServiceMasterSupplierImpl struct {
	MasterSupplierRepository repositorys.MasterSupplierRepository
}

func InitServiceMasterSupplier(unitRepository repositorys.MasterSupplierRepository) ServiceMasterSupplier {
	return &ServiceMasterSupplierImpl{
		MasterSupplierRepository: unitRepository,
	}
}

func (s *ServiceMasterSupplierImpl) ListMasterSupplierService() ([]*entities.MasterSupplier, schemas.SchemaDatabaseError) {

	res, err := s.MasterSupplierRepository.ListSupplierRepository()

	return res, err
}
