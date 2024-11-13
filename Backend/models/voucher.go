package models

import (
	"time"
)

type Voucher struct {
	Id           int        `json:"id" db:"id"`
	NamaVoucher  string     `json:"nama_voucher" db:"nama_voucher"`
	Status       string     `json:"status" db:"status"`
	Harga        int        `json:"harga" db:"harga"`
	Foto         string     `json:"foto" db:"foto"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
