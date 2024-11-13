package models

import "time"

type Manajemen struct {
	Id         int        `json:"id" db:"id"`
	Username   string     `json:"username" db:"username"`
	Password   string     `json:"password" db:"password"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	IdRole       int        `json:"id_role" db:"id_role"`
}
