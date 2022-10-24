package domain

import "disperinaker-api/model"

type SertRepoAdapter interface {
	CreateSert(sert model.Sertificate) (id int, err error)
	GetSerts() (serts []model.Sertificate, err error)
	GetSertByID(id int) (sert model.Sertificate, err error)
	UpdateSert(sert model.Sertificate, id int) error
	DeleteSert(id int) error
}
