package domain

import "disperinaker-api/model"

type SertRepoAdapter interface {
	CreateCalSert(sert model.CalonSertificate) (id int, err error)
	GetCalSerts() (serts []model.CalonSertificate, err error)
	GetCalSertByID(id int) (sert model.CalonSertificate, err error)
	UpdateCalSert(sert model.CalonSertificate, id int) error
	DeleteCalSert(id int) error
}
