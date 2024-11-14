package models

import "time"

type Hf struct {
	Id         int        `json:"id" db:"id"`
	Nama       string     `json:"nama" db:"nama"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}