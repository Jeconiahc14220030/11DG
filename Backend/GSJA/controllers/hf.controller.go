package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"strconv"

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

func GetAnggotaHF(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GETAnggotaHF(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAnggotaHF(idHF int) (models.Response, error) {
	var anggota models.Anggota
	var arrayAnggota []models.Anggota
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM anggota WHERE id_hf = ?"

	rows, err := con.Query(sqlStatement, idHF)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&anggota.Id,
			&anggota.IdHf,
			&anggota.Nama,
			&anggota.CreatedAt,
			&anggota.UpdatedAt,
			&anggota.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayAnggota = append(arrayAnggota, anggota)
	}

	response.Status = http.StatusOK
	response.Message = "Berhasil GET semua anggota HF"
	response.Data = arrayAnggota

	return response, err
}


func UpdateDeletedatHf(id int) (models.Response, error) {
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE hf SET deleted_at = NOW() WHERE id = ?"

	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return response, err
	}

	response.Status = http.StatusOK
	response.Message = "Berhasil DELETE HF"

	return response, err
}


func SoftDeleteHF(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := UpdateDeletedatHf(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}