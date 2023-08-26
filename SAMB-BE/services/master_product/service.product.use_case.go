package services

import (
	"SAMB-BE/entities"
	repositorys "SAMB-BE/repositorys/master_product"
	"SAMB-BE/schemas"
)

type ServiceMasterProduct interface {
	ListMasterProductService() ([]*entities.MasterProduct, schemas.SchemaDatabaseError)
}

type ServiceMasterProductImpl struct {
	MasterProductRepository repositorys.MasterProductRepository
}

func InitServiceMasterProduct(unitRepository repositorys.MasterProductRepository) ServiceMasterProduct {
	return &ServiceMasterProductImpl{
		MasterProductRepository: unitRepository,
	}
}

func (s *ServiceMasterProductImpl) ListMasterProductService() ([]*entities.MasterProduct, schemas.SchemaDatabaseError) {

	res, err := s.MasterProductRepository.ListProductRepository()

	return res, err
}
