package service

import (
	"disperinaker-api/model"
)

func (rs *sertService) CreatePesSertService(sert model.PesertaSertificate) (id int, err error) {
	return rs.repo.CreatePesSert(sert)
}

func (rs *sertService) GetPesSertsService() (serts []model.PesertaSertificate, err error) {
	return rs.repo.GetPesSerts()
}

func (rs *sertService) GetPesSertByIDService(id int) (sert model.PesertaSertificate, err error) {
	return rs.repo.GetPesSertByID(id)
}

func (rs *sertService) UpdatePesSertService(sert model.PesertaSertificate, id int) error {
	return rs.repo.UpdatePesSert(sert, id)
}

func (rs *sertService) DeletePesSertService(id int) error {
	return rs.repo.DeletePesSert(id)
}
