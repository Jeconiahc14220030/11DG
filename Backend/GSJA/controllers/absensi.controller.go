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
	idHfStr := c.FormValue("idHf")
	tanggal := c.FormValue("tanggal")

	idAnggota, err := strconv.Atoi(idAnggotaStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid idAnggota: " + idAnggotaStr})
	}

	idHf, err := strconv.Atoi(idHfStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid idHf: " + idHfStr})
	}

	absensiHf := models.AbsensiHf{
		IdAnggota: idAnggota,
		Idhf: idHf,
		Tanggal: tanggal,
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

	sqlStatement := "INSERT INTO absensi_hf (id_anggota, id_hf, tanggal) VALUES (?, ?, NOW())"
	_, err := con.Exec(sqlStatement, absensi.IdAnggota, absensi.IdJadwal)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Absensi hf berhasil ditambahkan"
	res.Data = absensi

	return res, nil
}