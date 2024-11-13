package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllJadwalLatihan(c echo.Context) error {
	result, err := GETJadwalLatihan()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETJadwalLatihan() (models.Response, error) {
	var jadwalLatihan models.JadwalLatihan
	var arrJadwalLatihan []models.JadwalLatihan
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM jadwal_latihan"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&jadwalLatihan.Id,
			&jadwalLatihan.Tanggal,
			&jadwalLatihan.Lokasi,
			&jadwalLatihan.CreatedAt,
			&jadwalLatihan.UpdatedAt,
			&jadwalLatihan.DeletedAt,
			&jadwalLatihan.IdAnggota,
			&jadwalLatihan.IdKomunitas,
		)

		if err != nil {
			return response, err
		}

		arrJadwalLatihan = append(arrJadwalLatihan, jadwalLatihan)
	} 

	response.Status = http.StatusOK
	response.Message = "Berhasil GET semua jadwal latihan"
	response.Data = arrJadwalLatihan

	return response, err
} 	