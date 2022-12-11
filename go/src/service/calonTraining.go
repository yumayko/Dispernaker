package service

import (
	"disperinaker-api/config"
	"disperinaker-api/domain"
	"disperinaker-api/model"
)

type trainService struct {
	conf config.Config
	repo domain.TrainRepoAdapter
}

func (rs *trainService) CreateCalTrainService(train model.PesertaTestMinatBakat) (id int, err error) {
	return rs.repo.CreateCalTrain(train)
}

func (rs *trainService) GetCalTrainsService() (trains []model.PesertaTestMinatBakat, err error) {
	return rs.repo.GetCalTrains()
}

func (rs *trainService) GetCalTrainByIDService(id int) (train model.PesertaTestMinatBakat, err error) {
	return rs.repo.GetCalTrainByID(id)
}

func (rs *trainService) UpdateCalTrainService(train model.PesertaTestMinatBakat, id int) error {
	return rs.repo.UpdateCalTrain(train, id)
}

func (rs *trainService) DeleteCalTrainService(id int) error {
	return rs.repo.DeleteCalTrain(id)
}

func NewTrainService(repo domain.TrainRepoAdapter, conf config.Config) domain.TrainServiceAdapter {
	return &trainService{
		repo: repo,
		conf: conf,
	}
}
