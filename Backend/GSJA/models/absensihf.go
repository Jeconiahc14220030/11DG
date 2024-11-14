package models

import "time"

type Absensihf struct {
	Id          int        `json:"id" db:"id"`
	Status      string     `json:"status" db:"status"`
	IdAnggota   int        `json:"id_anggota" db:"id_anggota"`
	IdJadwal    int        `json:"id_jadwal" db:"id_jadwal"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
