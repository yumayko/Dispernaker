package domain

import "disperinaker-api/model"

type SertServiceAdapter interface {
	CreateSertService(sert model.Sertificate) (id int, err error)
	GetSertsService() (serts []model.Sertificate, err error)
	GetSertByIDService(id int) (sert model.Sertificate, err error)
	UpdateSertService(sert model.Sertificate, id int) error
	DeleteSertService(id int) error
}
