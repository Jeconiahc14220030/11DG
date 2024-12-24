package models

import "time"

type Carousel struct {
	Id             int        `json:"id" db:"id"`
	Foto		string     `json:"foto" db:"foto"`
	Status		 string     `json:"status" db:"status"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
