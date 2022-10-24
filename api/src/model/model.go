package model

type Sertificate struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Kecamatan  string `json:"kecamatan"`
	Pelatihan  string `json:"pelatihan"`
	Keterangan string `json:"keterangan"`
}
