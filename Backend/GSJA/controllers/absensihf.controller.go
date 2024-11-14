package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllAbsensihf(c echo.Context) error {
	result, err := GETAllAbsensihf()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllAbsensihf() (models.Response, error) {
	var absensihf models.Absensihf
	var arrayAbsensihf []models.Absensihf
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM absensi_hf"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&absensihf.Id, 
			&absensihf.Status, 
			&absensihf.IdAnggota, 
			&absensihf.IdJadwal, 
			&absensihf.CreatedAt, 
			&absensihf.UpdatedAt, 
			&absensihf.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayAbsensihf = append(arrayAbsensihf, absensihf)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayAbsensihf

	return response, err
}

func FetchAbsensihfById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GETAbsensihfById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAbsensihfById(id int) (models.Response, error) {
	var absensihf models.Absensihf
	var arrayAbsensihf []models.Absensihf
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM absensi_hf WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&absensihf.Id, 
			&absensihf.Status, 
			&absensihf.IdAnggota, 
			&absensihf.IdJadwal, 
			&absensihf.CreatedAt, 
			&absensihf.UpdatedAt, 
			&absensihf.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayAbsensihf = append(arrayAbsensihf, absensihf)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayAbsensihf

	return response, err
}

