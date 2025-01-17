package models

import "time"

type AnggotaKomunitas struct {
	Id           int        `json:"id" db:"id"`
	IdAnggota    int        `json:"id_anggota" db:"id_anggota"`
	IdKomunitas  int        `json:"id_komunitas" db:"id_komunitas"`
	RequestAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	Status       string     `json:"status" db:"status"`
}
