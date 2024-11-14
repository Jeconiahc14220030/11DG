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
func AddBerita(c echo.Context) error {
	var berita models.Berita

	if err := c.Bind(&berita); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := POSTBerita(berita)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTBerita(berita models.Berita) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT INTO berita (deskripsi) VALUES (?)"
	_, err := con.Exec(sqlStatement, berita.Deskripsi)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Berita berhasil ditambahkan"
	res.Data = berita

	return res, nil
}