package repository

import (
	"disperinaker-api/model"
)

func (r *repoTrain) CreatePesTrain(train model.PesertaPelatihan) (id int, err error) {
	res := r.DB.Debug().Create(&train)
	if res.RowsAffected < 1 {
		return 0, res.Error
	}

	return train.ID, nil
}

func (r *repoTrain) GetPesTrains() (trains []model.PesertaPelatihan, err error) {
	if err = r.DB.
		Debug().
		Find(&trains).
		Error; err != nil {
		return []model.PesertaPelatihan{}, err
	}

	return trains, nil
}

func (r *repoTrain) GetPesTrainByID(id int) (train model.PesertaPelatihan, err error) {
	if err = r.DB.
		Debug().
		First(&train, id).
		Error; err != nil {
		return model.PesertaPelatihan{}, err
	}

	return
}

func (r *repoTrain) UpdatePesTrain(train model.PesertaPelatihan, id int) error {
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

func (r *repoTrain) DeletePesTrain(id int) error {
	train := model.PesertaPelatihan{}
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
