package models

import (
	"time"
)

type RenunganHarian struct {
	Id         int        `json:"id" db:"id"`
	Status     string     `json:"status" db:"status"`
	Isi        string     `json:"isi" db:"isi"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
