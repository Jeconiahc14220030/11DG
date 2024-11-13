package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllKutipanHarian(c echo.Context) error {
	result, err := GETAllKutipanHarian()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllKutipanHarian() (models.Response, error) {
	var kutipanHarian models.KutipanHarian
	var arrayKutipanHarian []models.KutipanHarian
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM kutipan_harian"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&kutipanHarian.Id,
			&kutipanHarian.Status,
			&kutipanHarian.Isi,
			&kutipanHarian.CreatedAt,
			&kutipanHarian.UpdatedAt,
			&kutipanHarian.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayKutipanHarian = append(arrayKutipanHarian, kutipanHarian)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayKutipanHarian

	return response, err
}