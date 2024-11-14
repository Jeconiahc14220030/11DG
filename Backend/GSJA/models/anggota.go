package models

import (
	_ "database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type Anggota struct {
	Id            int        `json:"id" db:"id"`
	Nama          string     `json:"nama" db:"nama"`
	Email         string     `json:"email" db:"email"`
	NomorTelepon  string     `json:"nomor_telepon" db:"nomor_telepon"`
	TanggalLahir  CustomDate  `json:"tanggal_lahir" db:"tanggal_lahir"`
	Poin          int        `json:"poin" db:"poin"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	IdKomunitas   int        `json:"id_komunitas" db:"id_komunitas"`
	IdHf          int        `json:"id_hf" db:"id_hf"`
}

type CustomDate struct {
    time.Time
}

const dateFormat = "2006-01-02"

func (c *CustomDate) UnmarshalJSON(b []byte) error {
    str := string(b)
    t, err := time.Parse(`"`+dateFormat+`"`, str)
    if err != nil {
        return err
    }
    c.Time = t
    return nil
}

func (c CustomDate) MarshalJSON() ([]byte, error) {
    return []byte(`"` + c.Time.Format(dateFormat) + `"`), nil
}

// To support database scanning
func (c *CustomDate) Value() (driver.Value, error) {
    return c.Time, nil
}

func (c *CustomDate) Scan(value interface{}) error {
    if t, ok := value.(time.Time); ok {
        c.Time = t
        return nil
    }
    return fmt.Errorf("unable to scan type %T into CustomDate", value)
}