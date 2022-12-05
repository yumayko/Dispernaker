package model

type CalonPesertaSertifikasi struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Nama       string `json:"nama"`
	Kecamatan  string `json:"kecamatan"`
	Pelatihan  string `json:"pelatihan"`
	Keterangan string `json:"keterangan"`
}

type PesertaSertifikasiSetelahPelatihan struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Nama       string `json:"nama"`
	Kecamatan  string `json:"kecamatan"`
	Pelatihan  string `json:"pelatihan"`
	Keterangan string `json:"keterangan"`
}

type PesertaTestMinatBakat struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Nama       string `json:"nama"`
	Kecamatan  string `json:"kecamatan"`
	Pelatihan  string `json:"pelatihan"`
	Keterangan string `json:"keterangan"`
}

type PesertaPelatihan struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Nama       string `json:"nama"`
	Kecamatan  string `json:"kecamatan"`
	Pelatihan  string `json:"pelatihan"`
	Keterangan string `json:"keterangan"`
}
