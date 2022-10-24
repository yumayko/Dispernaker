package domain

import "disperinaker-api/model"

type SertServiceAdapter interface {
	CreateSertService(sert model.Sertifikat) (id int, err error)
	GetSertsService() (serts []model.Sertifikat, err error)
	GetSertByIDService(id int) (sert model.Sertifikat, err error)
	UpdateSertService(sert model.Sertifikat, id int) error
	DeleteSertService(id int) error
}
