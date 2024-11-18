package models

import "time"

type Carousel struct {
	Id             int        `json:"id" db:"id"`
	Foto1          string     `json:"foto1" db:"foto1"`
	Foto2          string     `json:"foto2" db:"foto2"`
	Foto3          string     `json:"foto3" db:"foto3"`
	Foto4          string     `json:"foto4" db:"foto4"`
	StatusCarousel string     `json:"status_carousel" db:"status_carousel"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
