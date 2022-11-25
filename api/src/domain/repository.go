package domain

import "disperinaker-api/model"

type SertRepoAdapter interface {
	CreateCalSert(sert model.CalonSertificate) (id int, err error)
	GetCalSerts() (serts []model.CalonSertificate, err error)
	GetCalSertByID(id int) (sert model.CalonSertificate, err error)
	UpdateCalSert(sert model.CalonSertificate, id int) error
	DeleteCalSert(id int) error

	CreatePesSert(sert model.PesertaSertificate) (id int, err error)
	GetPesSerts() (serts []model.PesertaSertificate, err error)
	GetPesSertByID(id int) (sert model.PesertaSertificate, err error)
	UpdatePesSert(sert model.PesertaSertificate, id int) error
	DeletePesSert(id int) error
}
