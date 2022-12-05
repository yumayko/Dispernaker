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

func (rs *sertService) CreateCalSertService(sert model.CalonPesertaSertifikasi) (id int, err error) {
	return rs.repo.CreateCalSert(sert)
}

func (rs *sertService) GetCalSertsService() (serts []model.CalonPesertaSertifikasi, err error) {
	return rs.repo.GetCalSerts()
}

func (rs *sertService) GetCalSertByIDService(id int) (sert model.CalonPesertaSertifikasi, err error) {
	return rs.repo.GetCalSertByID(id)
}

func (rs *sertService) UpdateCalSertService(sert model.CalonPesertaSertifikasi, id int) error {
	return rs.repo.UpdateCalSert(sert, id)
}

func (rs *sertService) DeleteCalSertService(id int) error {
	return rs.repo.DeleteCalSert(id)
}

func NewSertService(repo domain.SertRepoAdapter, conf config.Config) domain.SertServiceAdapter {
	return &sertService{
		repo: repo,
		conf: conf,
	}
}
