package services

import (
	"SAMB-BE/entities"
	repositorys "SAMB-BE/repositorys/barang_masuk"
	"SAMB-BE/schemas"
	"time"
)

type ServicePenerimaanBarangImpl struct {
	PenerimaanBarangRepository repositorys.PenerimaanBarangRepository
}

func InitServicePenerimaanBarang(unitRepository repositorys.PenerimaanBarangRepository) ServicePenerimaanBarang {
	return &ServicePenerimaanBarangImpl{
		PenerimaanBarangRepository: unitRepository,
	}
}

func (s *ServicePenerimaanBarangImpl) CreatePenerimaanBarangHeaderService(input schemas.PenerimaanBarangHeaderRequest) (*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError) {

	Entity := entities.TrxPenerimaanBarangHeader{
		TrxInNo:      input.TrxInNo,
		WhsIdf:       input.WhsIdf,
		TrxInDate:    time.Now(),
		TrxInSuppIdf: input.TrxInSuppIdf,
		TrxInNotes:   input.TrxInNotes,
	}

	res, err := s.PenerimaanBarangRepository.CreatePenerimaanBarangHeaderRepository(Entity)

	return res, err
}

func (s *ServicePenerimaanBarangImpl) CreatePenerimaanBarangDetailService(input schemas.PenerimaanBarangDetailRequest) (*entities.TrxPenerimaanBarangDetail, schemas.SchemaDatabaseError) {

	Entity := entities.TrxPenerimaanBarangDetail{
		TrxInIDF:         input.TrxInIDF,
		TrxInDProductIdf: input.TrxInDProductIdf,
		TrxInDQtyDus:     input.TrxInDQtyDus,
		TrxInDQtyPcs:     input.TrxInDQtyPcs,
	}

	res, err := s.PenerimaanBarangRepository.CreatePenerimaanBarangDetailRepository(Entity)

	return res, err
}

func (s *ServicePenerimaanBarangImpl) ListBarangMasukHeaderService() ([]*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError) {

	res, err := s.PenerimaanBarangRepository.ListBarangMasukHeaderRepository()

	return res, err
}

func (s *ServicePenerimaanBarangImpl) ListBarangMasukDetailService(TrxInPK uint) ([]*entities.TrxPenerimaanBarangDetail, schemas.SchemaDatabaseError) {

	res, err := s.PenerimaanBarangRepository.ListBarangMasukDetailRepository(TrxInPK)

	return res, err
}

func (s *ServicePenerimaanBarangImpl) GetByIdBarangMasukHeaderService(TrxInPK uint) (*entities.TrxPenerimaanBarangHeader, schemas.SchemaDatabaseError) {

	res, err := s.PenerimaanBarangRepository.GetByIdBarangMasukHeaderRepository(TrxInPK)

	return res, err
}
