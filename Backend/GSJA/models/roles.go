package models

import (
	"time"
)

type Roles struct {
	Id         int        `json:"id" db:"id"`
	Roles       string     `json:"roles" db:"roles"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
