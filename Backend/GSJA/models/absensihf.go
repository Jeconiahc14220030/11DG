package models

import "time"

type AbsensiHf struct {
	Id        int        `json:"id" db:"id"`
	IdAnggota int        `json:"id_anggota" db:"id_anggota"`
	Idhf      int        `json:"idhf" db:"id_hf"`
	Topik    string     `json:"topik" db:"topik"`
	Tanggal   string     `json:"tanggal" db:"tanggal"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
