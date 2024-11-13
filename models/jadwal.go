package models

import "time"

type Jadwal struct {
	Id           int        `json:"id" db:"id"`
	Tanggal      time.Time  `json:"tanggal" db:"tanggal"`
	Topic        string     `json:"topik" db:"topik"`
	JenisIbadah  string     `json:"jenis_ibadah" db:"jenis_ibadah"`
	JumlahPoin   string     `json:"jumlah_poin" db:"jumlah_poin"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
