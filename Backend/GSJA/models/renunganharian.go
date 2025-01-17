package models

import (
	"time"
)

type RenunganHarian struct {
	Id         int        `json:"id" db:"id"`
	Isi        string     `json:"isi" db:"isi"`
	Foto 	 string     `json:"foto" db:"foto"`	
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
