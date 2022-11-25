package service

import (
	"disperinaker-api/config"
	"disperinaker-api/domain"
	"disperinaker-api/model"
)

type sertService struct {
	conf config.Config
	repo domain.SertRepoAdapter
}

func (rs *sertService) CreateSertService(sert model.CalonSertificate) (id int, err error) {
	return rs.repo.CreateSert(sert)
}

func (rs *sertService) GetSertsService() (serts []model.CalonSertificate, err error) {
	return rs.repo.GetSerts()
}

func (rs *sertService) GetSertByIDService(id int) (sert model.CalonSertificate, err error) {
	return rs.repo.GetSertByID(id)
}

func (rs *sertService) UpdateSertService(sert model.CalonSertificate, id int) error {
	return rs.repo.UpdateSert(sert, id)
}

func (rs *sertService) DeleteSertService(id int) error {
	return rs.repo.DeleteSert(id)
}

func NewSertService(repo domain.SertRepoAdapter, conf config.Config) domain.SertServiceAdapter {
	return &sertService{
		repo: repo,
		conf: conf,
	}
}
