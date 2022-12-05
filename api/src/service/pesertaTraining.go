package service

import (
	"disperinaker-api/model"
)

func (rs *trainService) CreatePesTrainService(train model.PesertaPelatihan) (id int, err error) {
	return rs.repo.CreatePesTrain(train)
}

func (rs *trainService) GetPesTrainsService() (trains []model.PesertaPelatihan, err error) {
	return rs.repo.GetPesTrains()
}

func (rs *trainService) GetPesTrainByIDService(id int) (train model.PesertaPelatihan, err error) {
	return rs.repo.GetPesTrainByID(id)
}

func (rs *trainService) UpdatePesTrainService(train model.PesertaPelatihan, id int) error {
	return rs.repo.UpdatePesTrain(train, id)
}

func (rs *trainService) DeletePesTrainService(id int) error {
	return rs.repo.DeletePesTrain(id)
}
