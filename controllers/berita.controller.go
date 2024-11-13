package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllBerita(c echo.Context) error {
	result, err := GETAllBerita()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllBerita() (models.Response, error) {
	var berita models.Berita
	var arrBerita []models.Berita
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM berita"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&berita.Id,
			&berita.Deskripsi,
			&berita.CreatedAt,
			&berita.UpdatedAt,
			&berita.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrBerita = append(arrBerita, berita)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrBerita
	
	return response, err
}