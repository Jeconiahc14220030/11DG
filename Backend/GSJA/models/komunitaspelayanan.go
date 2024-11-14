package models

import (
	"time"
)

type KomunitasPelayanan struct {
	Id             int        `json:"id" db:"id"`
	NamaKomunitas  string     `json:"nama_komunitas" db:"nama_komunitas"`
	Deskripsi      string     `json:"deskripsi" db:"deskripsi"`
	Snk            string     `json:"snk" db:"snk"`
	Manfaat        string     `json:"manfaat" db:"manfaat"`
	Gambar         string     `json:"gambar" db:"gambar"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
