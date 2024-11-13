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
			&renungan_harian.Status,
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