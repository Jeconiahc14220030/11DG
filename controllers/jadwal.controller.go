package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllJadwal(c echo.Context) error {
	result, err := GETAllJadwal()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllJadwal() (models.Response, error) {
	var jadwal models.Jadwal
	var arrJadwal []models.Jadwal
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM jadwal"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&jadwal.Id, 
			&jadwal.Tanggal, 
			&jadwal.Topic, 
			&jadwal.JenisIbadah, 
			&jadwal.JumlahPoin,
			&jadwal.CreatedAt,
			&jadwal.UpdatedAt,
			&jadwal.DeletedAt,
		)

		if err != nil {
			return res, err
		}

		arrJadwal = append(arrJadwal, jadwal)
	}

	res.Status = http.StatusOK
	res.Message = "OK"
	res.Data = arrJadwal

	return res, err
}
