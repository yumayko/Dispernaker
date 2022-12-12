package service

import (
	"disperinaker-api/model"
)

func (rs *sertService) CreatePesSertService(sert model.PesertaSertifikasiSetelahPelatihan) (id int, err error) {
	return rs.repo.CreatePesSert(sert)
}

func (rs *sertService) GetPesSertsService() (serts []model.PesertaSertifikasiSetelahPelatihan, err error) {
	return rs.repo.GetPesSerts()
}

func (rs *sertService) GetPesSertByIDService(id int) (sert model.PesertaSertifikasiSetelahPelatihan, err error) {
	return rs.repo.GetPesSertByID(id)
}

func (rs *sertService) UpdatePesSertService(sert model.PesertaSertifikasiSetelahPelatihan, id int) error {
	return rs.repo.UpdatePesSert(sert, id)
}

func (rs *sertService) DeletePesSertService(id int) error {
	return rs.repo.DeletePesSert(id)
}
