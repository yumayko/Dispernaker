package repository

import (
	"disperinaker-api/domain"
	"disperinaker-api/model"

	"gorm.io/gorm"
)

type repoTrain struct {
	DB *gorm.DB
}

func (r *repoTrain) CreateCalTrain(train model.PesertaTestMinatBakat) (id int, err error) {
	res := r.DB.Debug().Create(&train)
	if res.RowsAffected < 1 {
		return 0, res.Error
	}

	return train.ID, nil
}

func (r *repoTrain) GetCalTrains() (trains []model.PesertaTestMinatBakat, err error) {
	if err = r.DB.
		Debug().
		Find(&trains).
		Error; err != nil {
		return []model.PesertaTestMinatBakat{}, err
	}

	return trains, nil
}

func (r *repoTrain) GetCalTrainByID(id int) (train model.PesertaTestMinatBakat, err error) {
	if err = r.DB.
		Debug().
		First(&train, id).
		Error; err != nil {
		return model.PesertaTestMinatBakat{}, err
	}

	return
}

func (r *repoTrain) UpdateCalTrain(train model.PesertaTestMinatBakat, id int) error {
	train.ID = id

	res := r.DB.
		Debug().
		Where("id = ?", id).
		UpdateColumns(&train)
	if res.RowsAffected < 1 {
		return res.Error
	}

	return nil
}

func (r *repoTrain) DeleteCalTrain(id int) error {
	train := model.PesertaTestMinatBakat{}
	train.ID = id

	res := r.DB.Find(&train)

	if res.RowsAffected < 1 {
		return res.Error
	}

	if err := r.DB.
		Debug().
		Unscoped().
		Delete(&train).
		Error; err != nil {
		return err
	}

	return nil
}

func NewTrain(db *gorm.DB) domain.TrainRepoAdapter {
	return &repoTrain{
		DB: db,
	}
}
