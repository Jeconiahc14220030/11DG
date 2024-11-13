package models

import (
	"time"
)

type RolesAnggota struct {
	Id        int        `json:"id" db:"id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at" db:"deleted_at"`
	IdAnggota int        `json:"id_anggota" db:"id_anggota"`
	IdRole    int        `json:"id_role" db:"id_role"`
}