package models

import (
	_ "database/sql"
	"time"
)

type Anggota struct {
	Id           int        `json:"id" db:"id"`
	Nama         string     `json:"nama" db:"nama"`
	Username     string     `json:"username" db:"username"`
	Password     string     `json:"password" db"password"`
	Email        string     `json:"email" db:"email"`
	NomorTelepon string     `json:"nomor_telepon" db:"nomor_telepon"`
	TanggalLahir string     `json:"tanggal_lahir" db:"tanggal_lahir"`
	Poin         int        `json:"poin" db:"poin"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	IdKomunitas  int        `json:"id_komunitas" db:"id_komunitas"`
	IdHf         int        `json:"id_hf" db:"id_hf"`
}
