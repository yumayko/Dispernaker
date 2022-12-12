package domain

import "disperinaker-api/model"

type SertRepoAdapter interface {
	CreateCalSert(sert model.CalonPesertaSertifikasi) (id int, err error)
	GetCalSerts() (serts []model.CalonPesertaSertifikasi, err error)
	GetCalSertByID(id int) (sert model.CalonPesertaSertifikasi, err error)
	UpdateCalSert(sert model.CalonPesertaSertifikasi, id int) error
	DeleteCalSert(id int) error

	CreatePesSert(sert model.PesertaSertifikasiSetelahPelatihan) (id int, err error)
	GetPesSerts() (serts []model.PesertaSertifikasiSetelahPelatihan, err error)
	GetPesSertByID(id int) (sert model.PesertaSertifikasiSetelahPelatihan, err error)
	UpdatePesSert(sert model.PesertaSertifikasiSetelahPelatihan, id int) error
	DeletePesSert(id int) error
}

type TrainRepoAdapter interface {
	CreateCalTrain(train model.PesertaTestMinatBakat) (id int, err error)
	GetCalTrains() (trains []model.PesertaTestMinatBakat, err error)
	GetCalTrainByID(id int) (train model.PesertaTestMinatBakat, err error)
	UpdateCalTrain(train model.PesertaTestMinatBakat, id int) error
	DeleteCalTrain(id int) error

	CreatePesTrain(train model.PesertaPelatihan) (id int, err error)
	GetPesTrains() (trains []model.PesertaPelatihan, err error)
	GetPesTrainByID(id int) (train model.PesertaPelatihan, err error)
	UpdatePesTrain(train model.PesertaPelatihan, id int) error
	DeletePesTrain(id int) error
}
