package models

import "time"

type Pengumuman struct {
	Id           int       `json:"id" db:"id"`
	Konten       string    `json:"konten" db:"konten"`
	Tanggal      time.Time `json:"tanggal" db:"tanggal"`
	Created_at   time.Time `json:"created_at" db:"created_at"`
	Updated_at   time.Time `json:"updated_at" db:"updated_at"`
	Deleted_at   time.Time `json:"deleted_at" db:"deleted_at"`
	Id_komunitas int       `json:"id_komunitas" db:"id_komunitas"`
}
