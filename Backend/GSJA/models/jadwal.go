package models

import "time"

type Jadwal struct {
	Id           int        `json:"id" db:"id"`
	Tanggal      string  `json:"tanggal" db:"tanggal"`
	Topik        string     `json:"topik" db:"topik"`
	JenisIbadah  string     `json:"jenis_ibadah" db:"jenis_ibadah"`
	JumlahPoin   int     `json:"jumlah_poin" db:"jumlah_poin"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
