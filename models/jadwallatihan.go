package models

import "time"

type JadwalLatihan struct {
	Id           int        `json:"id" db:"id"`
	Tanggal      time.Time  `json:"tanggal" db:"tanggal"`
	Lokasi       string     `json:"lokasi" db:"lokasi"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	IdAnggota    int        `json:"id_anggota" db:"id_anggota"`
	IdKomunitas  int        `json:"id_komunitas" db:"id_komunitas"`
}
