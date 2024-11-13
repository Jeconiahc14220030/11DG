package models

import "time"

type Berita struct {
	Id         int        `json:"id" db:"id"`
	Deskripsi  string     `json:"deskripsi" db:"deskripsi"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
