package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllHf(c echo.Context) error {
	result, err := GETAllHf()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllHf() (models.Response, error) {
	var hf models.Hf
	var arrayHf []models.Hf
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM hf"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&hf.Id,
			&hf.Nama,
			&hf.CreatedAt,
			&hf.UpdatedAt,
			&hf.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayHf = append(arrayHf, hf)
	}

	response.Status = http.StatusOK
	response.Message = "Berhasil GET semua HF"
	response.Data = arrayHf

	return response, err
}
