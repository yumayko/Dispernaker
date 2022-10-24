package domain

import "disperinaker-api/model"

type SertRepoAdapter interface {
	CreateSert(sert model.Sertifikat) (id int, err error)
	GetSerts() (serts []model.Sertifikat, err error)
	GetSertByID(id int) (sert model.Sertifikat, err error)
	UpdateSert(sert model.Sertifikat, id int) error
	DeleteSert(id int) error
}
