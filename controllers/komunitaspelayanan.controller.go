package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllKomunitas(c echo.Context) error {
	result, err := GETAllKomunitas()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllKomunitas() (models.Response, error) {
	var komunitas models.KomunitasPelayanan
	var arrayKomunitas []models.KomunitasPelayanan
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM komunitas_pelayanan"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&komunitas.Id, 
			&komunitas.NamaKomunitas, 
			&komunitas.Deskripsi, 
			&komunitas.Snk, 
			&komunitas.Manfaat, 
			&komunitas.Gambar,
			&komunitas.CreatedAt,
			&komunitas.UpdatedAt,
			&komunitas.DeletedAt,
		)

		if err != nil {
			return response, nil
		}

		arrayKomunitas = append(arrayKomunitas, komunitas)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayKomunitas

	return response, err
}

func FetchKomunitasById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GETKomunitasById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETKomunitasById(id int) (models.Response, error) {
	var komunitas models.KomunitasPelayanan
	var arrayKomunitas []models.KomunitasPelayanan
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM komunitas_pelayanan WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&komunitas.Id, 
			&komunitas.NamaKomunitas, 
			&komunitas.Deskripsi, 
			&komunitas.Snk, 
			&komunitas.Manfaat, 
			&komunitas.Gambar,
			&komunitas.CreatedAt,
			&komunitas.UpdatedAt,
			&komunitas.DeletedAt,
		)

		if err != nil {
			return response, nil
		}

		arrayKomunitas = append(arrayKomunitas, komunitas)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayKomunitas

	return response, err
}
