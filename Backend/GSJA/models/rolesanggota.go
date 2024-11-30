package models

import (
	"time"
)

type RolesAnggota struct {
	Id        int        `json:"id" db:"id"`
	IdAnggota int        `json:"id_anggota" db:"id_anggota"`
	IdRoles    int        `json:"id_roles" db:"id_roles"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at" db:"deleted_at"`

}