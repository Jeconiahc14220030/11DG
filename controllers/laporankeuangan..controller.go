package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllLaporanKeuangan(c echo.Context) error {
	result, err := GETLaporanKeuangan()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETLaporanKeuangan() (models.Response, error) {
	var laporan models.LaporanKeuangan
	var arrLaporan []models.LaporanKeuangan
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM laporankeuangan "

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&laporan.Id, 
			&laporan.Tanggal, 
			&laporan.Jenis, 
			&laporan.Nominal, 
			&laporan.CreatedAt, 
			&laporan.UpdatedAt, 
			&laporan.DeletedAt, 
			&laporan.IdPembuat,
		)

		if err != nil {
			return response, err
		}

		arrLaporan = append(arrLaporan, laporan)
	}

	response.Status = http.StatusOK
	response.Message = "Berhasil GET semua laporan keuangan"
	response.Data = arrLaporan

	return response, err
}

func FetchLaporanKeuanganById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GETLaporanKeuanganById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETLaporanKeuanganById(id int) (models.Response, error) {
	var laporan models.LaporanKeuangan
	var arrLaporan []models.LaporanKeuangan
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM laporankeuangan WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&laporan.Id, 
			&laporan.Tanggal, 
			&laporan.Jenis, 
			&laporan.Nominal, 
			&laporan.CreatedAt, 
			&laporan.UpdatedAt, 
			&laporan.DeletedAt, 
			&laporan.IdPembuat,
		)

		if err != nil {
			return response, err
		}

		arrLaporan = append(arrLaporan, laporan)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrLaporan

	return response, err
}
