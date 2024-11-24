package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllRenunganHarian(c echo.Context) error {
	result, err := GETAllRenunganHarian()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllRenunganHarian() (models.Response, error) {
	var renungan_harian models.RenunganHarian
	var arrayRenungan []models.RenunganHarian
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM renungan_harian"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&renungan_harian.Id,
			// &renungan_harian.Status,
			&renungan_harian.Isi,
			&renungan_harian.CreatedAt,
			&renungan_harian.UpdatedAt,
			&renungan_harian.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayRenungan = append(arrayRenungan, renungan_harian)

	}
	
	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayRenungan

	return response, err
}

func AddRenunganHarian(c echo.Context) error {
	status := c.FormValue("status")
	isi := c.FormValue("isi")
	
	renunganHarian := models.RenunganHarian{
		Status: status,
		Isi: isi,
	}

	if err := c.Bind(&renunganHarian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := POSTRenunganHarian(renunganHarian)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTRenunganHarian(renunganHarian models.RenunganHarian) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT INTO renungan_harian (status, isi) VALUES (?, ?)"

	_, err := con.Exec(sqlStatement, renunganHarian.Status, renunganHarian.Isi)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Renungan harian berhasil ditambahkan"
	res.Data = renunganHarian

	return res, nil
}

func SoftDeletedataRenunganHarian(c echo.Context) error {
	id := c.Param("id")

	result, err := UpdateDeletedAtRenunganHarian(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtRenunganHarian(id string) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "UPDATE renungan_harian SET deleted_at = NOW() WHERE id = ?"

	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Renungan harian berhasil dihapus"

	return res, nil
}
