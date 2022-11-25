package domain

import "disperinaker-api/model"

type SertRepoAdapter interface {
	CreateSert(sert model.CalonSertificate) (id int, err error)
	GetSerts() (serts []model.CalonSertificate, err error)
	GetSertByID(id int) (sert model.CalonSertificate, err error)
	UpdateSert(sert model.CalonSertificate, id int) error
	DeleteSert(id int) error
}
