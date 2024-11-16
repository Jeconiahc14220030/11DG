package models

import (
	_ "database/sql"
	"time"
)

type Requestkomunitas struct {
	Id           int        `json:"id" db:"id_request"`
	Id_komunitas int        `json:"id_komunitas" db:"id_komunitas"`
	Id_anggota   int        `json:"id_anggota" db:"id_anggota"`
	RequestAt    time.Time  `json:"request_at" db:"request_at"`
	UpdatedAt    *time.Time  `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"delete_at"`
	Status       string     `json:"status" db:"status"`
}
