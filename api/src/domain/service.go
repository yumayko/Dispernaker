package domain

import "disperinaker-api/model"

type SertServiceAdapter interface {
	CreateCalSertService(sert model.CalonSertificate) (id int, err error)
	GetCalSertsService() (serts []model.CalonSertificate, err error)
	GetCalSertByIDService(id int) (sert model.CalonSertificate, err error)
	UpdateCalSertService(sert model.CalonSertificate, id int) error
	DeleteCalSertService(id int) error
}
