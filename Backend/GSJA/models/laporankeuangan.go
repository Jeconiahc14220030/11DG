package models

import "time"

type LaporanKeuangan struct {
	Id         int        `json:"id" db:"id"`
	Tanggal    string     `json:"tanggal" db:"tanggal"`
	Jenis      string     `json:"jenis" db:"jenis"`
	Nominal    string     `json:"nominal" db:"nominal"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	IdPembuat  int        `json:"id_pembuat" db:"id_pembuat"`
}