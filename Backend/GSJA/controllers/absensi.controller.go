package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddKehadiranAbsensi(c echo.Context) error {
	idAnggotaStr := c.FormValue("idAnggota")
	idJadwalStr := c.FormValue("idJadwal")

	idJadwal, err := strconv.Atoi(idJadwalStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid idJadwal" + idJadwalStr})
	}

	idAnggota, err := strconv.Atoi(idAnggotaStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid idAnggota: " + idAnggotaStr})
	}

	absensiHf := models.AbsensiHf{
		Status:    "hadir",
		IdAnggota: idAnggota,
		IdJadwal:  idJadwal,
	}

	result, err := POSTAbsensiHf(absensiHf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTAbsensi(absensi models.Absensi) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT INTO absensi_hf (status, id_anggota, id_jadwal) VALUES (?, ?, ?)"
	_, err := con.Exec(sqlStatement, absensi.Status, absensi.IdAnggota, absensi.IdJadwal)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Absensi hf berhasil ditambahkan"
	res.Data = absensi

	return res, nil
}