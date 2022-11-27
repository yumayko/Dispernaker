package domain

import "disperinaker-api/model"

type SertServiceAdapter interface {
	CreateCalSertService(sert model.CalonSertificate) (id int, err error)
	GetCalSertsService() (serts []model.CalonSertificate, err error)
	GetCalSertByIDService(id int) (sert model.CalonSertificate, err error)
	UpdateCalSertService(sert model.CalonSertificate, id int) error
	DeleteCalSertService(id int) error

	CreatePesSertService(sert model.PesertaSertificate) (id int, err error)
	GetPesSertsService() (serts []model.PesertaSertificate, err error)
	GetPesSertByIDService(id int) (sert model.PesertaSertificate, err error)
	UpdatePesSertService(sert model.PesertaSertificate, id int) error
	DeletePesSertService(id int) error
}

type TrainServiceAdapter interface {
	CreateCalTrainService(train model.CalonTraining) (id int, err error)
	GetCalTrainsService() (trains []model.CalonTraining, err error)
	GetCalTrainByIDService(id int) (train model.CalonTraining, err error)
	UpdateCalTrainService(train model.CalonTraining, id int) error
	DeleteCalTrainService(id int) error

	CreatePesTrainService(train model.PesertaTraining) (id int, err error)
	GetPesTrainsService() (trains []model.PesertaTraining, err error)
	GetPesTrainByIDService(id int) (train model.PesertaTraining, err error)
	UpdatePesTrainService(train model.PesertaTraining, id int) error
	DeletePesTrainService(id int) error
}
