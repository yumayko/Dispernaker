package repository

import (
	"disperinaker-api/domain"
	"disperinaker-api/model"

	"gorm.io/gorm"
)

type repoSert struct {
	DB *gorm.DB
}

func (r *repoSert) CreateSert(sert model.Sertifikat) (id int, err error) {
	res := r.DB.Debug().Create(&sert)
	if res.RowsAffected < 1 {
		return 0, res.Error
	}

	return sert.ID, nil
}

func (r *repoSert) GetSerts() (serts []model.Sertifikat, err error) {
	if err = r.DB.
		Debug().
		Find(&serts).
		Error; err != nil {
		return []model.Sertifikat{}, err
	}

	return serts, nil
}

func (r *repoSert) GetSertByID(id int) (sert model.Sertifikat, err error) {
	if err = r.DB.
		Debug().
		First(&sert, id).
		Error; err != nil {
		return model.Sertifikat{}, err
	}

	return
}

func (r *repoSert) UpdateSert(sert model.Sertifikat, id int) error {
	temp := model.Sertifikat{}
	temp.ID = id

	res := r.DB.
		Debug().
		Save(&temp)
	if res.RowsAffected < 1 {
		return res.Error
	}

	return nil
}

func (r *repoSert) DeleteSert(id int) error {
	sert := model.Sertifikat{}
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
