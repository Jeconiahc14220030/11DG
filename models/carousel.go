package models

import "time"

type Carousel struct {
	Id             int        `json:"id" db:"id"`
	Foto1          string     `json:"foto_1" db:"foto_1"`
	Foto2          string     `json:"foto_2" db:"foto_2"`
	Foto3          string     `json:"foto_3" db:"foto_3"`
	Foto4          string     `json:"foto_4" db:"foto_4"`
	StatusCarousel int        `json:"status_carousel" db:"status_carousel"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
