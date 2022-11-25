package domain

import "disperinaker-api/model"

type SertServiceAdapter interface {
	CreateSertService(sert model.CalonSertificate) (id int, err error)
	GetSertsService() (serts []model.CalonSertificate, err error)
	GetSertByIDService(id int) (sert model.CalonSertificate, err error)
	UpdateSertService(sert model.CalonSertificate, id int) error
	DeleteSertService(id int) error
}
