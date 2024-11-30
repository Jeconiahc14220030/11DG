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

	sqlStatement := "SELECT * FROM laporan_keuangan"

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

	sqlStatement := "SELECT * FROM laporan_keuangan WHERE id = ?"

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

func AddLaporanKeuangan(c echo.Context) error {
	tanggal := c.FormValue("tanggal")
	jenis := c.FormValue("jenis")
	nominal := c.FormValue("nominal")
	strIdPembuat := c.FormValue("id_pembuat")

	idPembuat, err := strconv.Atoi(strIdPembuat)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id pembuat"})
	}

	laporanKeuangan := models.LaporanKeuangan{
		Tanggal: tanggal,
		Jenis: jenis,
		Nominal: nominal,
		IdPembuat: idPembuat,
	}

	if err := c.Bind(&laporanKeuangan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := POSTLaporanKeuangan(laporanKeuangan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTLaporanKeuangan(laporanKeuangan models.LaporanKeuangan) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT INTO laporan_keuangan (tanggal, jenis, nominal, id_pembuat) VALUES (?, ?, ?, ?)"

	_, err := con.Exec(sqlStatement, laporanKeuangan.Tanggal, laporanKeuangan.Jenis, laporanKeuangan.Nominal, laporanKeuangan.IdPembuat)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Laporan keuangan berhasil ditambahkan"
	res.Data = laporanKeuangan

	return res, nil
}

func SoftDeletedataLaporanKeuangan(c echo.Context) error {
	id := c.Param("id")

	result, err := UpdateDeletedAtLaporanKeuangan(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtLaporanKeuangan(id string) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "UPDATE laporan_keuangan SET deleted_at = NOW() WHERE id = ?"

	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Laporan keuangan berhasil dihapus"

	return res, nil
}