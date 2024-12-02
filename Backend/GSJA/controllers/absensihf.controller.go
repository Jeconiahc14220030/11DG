package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"fmt"
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
	var absensihf models.AbsensiHf
	var arrayAbsensihf []models.AbsensiHf
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
			&absensihf.IdAnggota, 
			&absensihf.Idhf,
			&absensihf.Topik,
			&absensihf.Tanggal,
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
	var absensihf models.AbsensiHf
	var arrayAbsensihf []models.AbsensiHf
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM absensi_hf WHERE id_hf = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&absensihf.Id,
			&absensihf.IdAnggota, 
			&absensihf.Idhf,
			&absensihf.Topik,
			&absensihf.Tanggal,
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

func FetchAbsensihfByAnggotaId(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GetAbsensiHfAnggotaById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAbsensiHfAnggotaById(id_anggota int) (models.Response, error) {
	var absensihf models.AbsensiHf
	var arrayAbsensihf []models.AbsensiHf
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM absensi_hf WHERE id_anggota = ?"

	rows, err := con.Query(sqlStatement, id_anggota)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&absensihf.Id,
			&absensihf.IdAnggota,
			&absensihf.Idhf,
			&absensihf.Topik,
			&absensihf.Tanggal,
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

func AddAbsensiHfForm(c echo.Context) error {
	idAnggotaStr := c.FormValue("id_anggota")
	idHFStr := c.FormValue("id_hf")
	topik := c.FormValue("topik")
	tanggal := c.FormValue("tanggal")
	idAnggota, err := strconv.Atoi(idAnggotaStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid idAnggota: " + idAnggotaStr})
	}

	idHF, err := strconv.Atoi(idHFStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id_hf"})
	}

	absensiHf := models.AbsensiHf{
		IdAnggota: idAnggota,
		Idhf: idHF,
		Topik: topik,
		Tanggal: tanggal,
	}

	result, err := POSTAbsensiHf(absensiHf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTAbsensiHf(absensiHf models.AbsensiHf) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO absensi_hf (id_anggota, id_hf, tanggal) VALUES (?, ?, ?)"
	_, err := con.Exec(sqlStatement, absensiHf.IdAnggota, absensiHf.Idhf, absensiHf.Tanggal)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Absensi hf berhasil ditambahkan"
	res.Data = absensiHf

	return res, nil
}

func AddAbsensiHfJson(c echo.Context) error {
	AbsensiHf := models.AbsensiHf{}
	if err := c.Bind(&AbsensiHf); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	con := db.CreateCon()

	sqlStatement := "INSERT INTO absensi_hf (id_anggota, id_hf, topik, tanggal) VALUES (?, ?, ?, ?)"
	_, err := con.Exec(sqlStatement, AbsensiHf.IdAnggota, AbsensiHf.Idhf, AbsensiHf.Topik, AbsensiHf.Tanggal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, AbsensiHf)
}

func SoftDeleteAbsensiHf(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := UpdateDeletedatAbsensiHf(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedatAbsensiHf(id int) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	
	sqlStatement := "UPDATE absensi_hf SET deleted_at = NOW() WHERE id = ?"
	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Deleted at updated successfully"
	res.Data = map[string]int{"id": id}

	return res, nil
}