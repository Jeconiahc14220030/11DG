package models

import "time"

type KontenGereja struct {
	Id                int        `json:"id" db:"id"`
	Visi              string     `json:"visi" db:"visi"`
	Misi              string     `json:"misi" db:"misi"`
	Tujuan            string     `json:"tujuan" db:"tujuan"`
	PesanKetua        string     `json:"pesan_ketua" db:"pesan_ketua"`
	IdCarousel        int        `json:"id_carousel" db:"id_carousel"`
	IdBerita          string     `json:"id_berita" db:"id_berita"`
	IdKutipanHarian   int        `json:"id_kutipan_harian" db:"id_kutipan_harian"`
	IdRenunganHarian  int        `json:"id_renungan_harian" db:"id_renungan_harian"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
