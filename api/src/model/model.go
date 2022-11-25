package model

type CalonSertificate struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Kecamatan  string `json:"kecamatan"`
	Pelatihan  string `json:"pelatihan"`
	Keterangan string `json:"keterangan"`
}

type PesertaSertificate struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Kecamatan  string `json:"kecamatan"`
	Pelatihan  string `json:"pelatihan"`
	Keterangan string `json:"keterangan"`
}

type CalonTraining struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Kecamatan  string `json:"kecamatan"`
	Pelatihan  string `json:"pelatihan"`
	Keterangan string `json:"keterangan"`
}

type PesertaTraining struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Kecamatan  string `json:"kecamatan"`
	Pelatihan  string `json:"pelatihan"`
	Keterangan string `json:"keterangan"`
}
