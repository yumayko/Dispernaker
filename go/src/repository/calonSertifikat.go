package repository

import (
	"disperinaker-api/domain"
	"disperinaker-api/model"

	"gorm.io/gorm"
)

type repoSert struct {
	DB *gorm.DB
}

func (r *repoSert) CreateCalSert(sert model.CalonPesertaSertifikasi) (id int, err error) {
	res := r.DB.Debug().Create(&sert)
	if res.RowsAffected < 1 {
		return 0, res.Error
	}

	return sert.ID, nil
}

func (r *repoSert) GetCalSerts() (serts []model.CalonPesertaSertifikasi, err error) {
	if err = r.DB.
		Debug().
		Find(&serts).
		Error; err != nil {
		return []model.CalonPesertaSertifikasi{}, err
	}

	return serts, nil
}

func (r *repoSert) GetCalSertByID(id int) (sert model.CalonPesertaSertifikasi, err error) {
	if err = r.DB.
		Debug().
		First(&sert, id).
		Error; err != nil {
		return model.CalonPesertaSertifikasi{}, err
	}

	return
}

func (r *repoSert) UpdateCalSert(sert model.CalonPesertaSertifikasi, id int) error {
	sert.ID = id

	res := r.DB.
		Debug().
		Where("id = ?", id).
		UpdateColumns(&sert)
	if res.RowsAffected < 1 {
		return res.Error
	}

	return nil
}

func (r *repoSert) DeleteCalSert(id int) error {
	sert := model.CalonPesertaSertifikasi{}
	sert.ID = id

	res := r.DB.Find(&sert)

	if res.RowsAffected < 1 {
		return res.Error
	}

	if err := r.DB.
		Debug().
		Unscoped().
		Delete(&sert).
		Error; err != nil {
		return err
	}

	return nil
}

func NewSert(db *gorm.DB) domain.SertRepoAdapter {
	return &repoSert{
		DB: db,
	}
}
