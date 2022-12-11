package domain

import "disperinaker-api/model"

type SertServiceAdapter interface {
	CreateCalSertService(sert model.CalonPesertaSertifikasi) (id int, err error)
	GetCalSertsService() (serts []model.CalonPesertaSertifikasi, err error)
	GetCalSertByIDService(id int) (sert model.CalonPesertaSertifikasi, err error)
	UpdateCalSertService(sert model.CalonPesertaSertifikasi, id int) error
	DeleteCalSertService(id int) error

	CreatePesSertService(sert model.PesertaSertifikasiSetelahPelatihan) (id int, err error)
	GetPesSertsService() (serts []model.PesertaSertifikasiSetelahPelatihan, err error)
	GetPesSertByIDService(id int) (sert model.PesertaSertifikasiSetelahPelatihan, err error)
	UpdatePesSertService(sert model.PesertaSertifikasiSetelahPelatihan, id int) error
	DeletePesSertService(id int) error
}

type TrainServiceAdapter interface {
	CreateCalTrainService(train model.PesertaTestMinatBakat) (id int, err error)
	GetCalTrainsService() (trains []model.PesertaTestMinatBakat, err error)
	GetCalTrainByIDService(id int) (train model.PesertaTestMinatBakat, err error)
	UpdateCalTrainService(train model.PesertaTestMinatBakat, id int) error
	DeleteCalTrainService(id int) error

	CreatePesTrainService(train model.PesertaPelatihan) (id int, err error)
	GetPesTrainsService() (trains []model.PesertaPelatihan, err error)
	GetPesTrainByIDService(id int) (train model.PesertaPelatihan, err error)
	UpdatePesTrainService(train model.PesertaPelatihan, id int) error
	DeletePesTrainService(id int) error
}
