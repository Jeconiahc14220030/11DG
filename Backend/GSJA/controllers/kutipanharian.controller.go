package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"strconv"

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

func AddKutipanHarian(c echo.Context) error {
	isi := c.FormValue("isi")

	kutipanHarian := models.KutipanHarian{
		Isi: isi,
	}

	if err := c.Bind(&kutipanHarian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := POSTKutipanHarian(kutipanHarian)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTKutipanHarian(kutipanHarian models.KutipanHarian) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO kutipan_harian (isi) VALUES (?)"
	_, err := con.Exec(sqlStatement, kutipanHarian.Isi)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Kutipan harian berhasil ditambahkan"
	res.Data = kutipanHarian

	return res, nil
}

func SoftDeletedataKutipanHarian(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id"})
	}

	result, err := UpdateDeletedAtKutipanHarian(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtKutipanHarian(id int) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE kutipan_harian SET deleted_at = NOW() WHERE id = ?"
	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Kutipan harian berhasil dihapus"

	return res, nil
}