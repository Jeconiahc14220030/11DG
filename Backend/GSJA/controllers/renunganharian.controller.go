package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

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
			&renungan_harian.Isi,
			&renungan_harian.Foto,
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

func AddRenunganHarian(c echo.Context) error {
	isi := c.FormValue("isi")
	foto, err := c.FormFile("foto")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	namaFoto := foto.Filename
	
	renunganHarian := models.RenunganHarian{
		Isi: isi,
		Foto: namaFoto,
	}

	if err := c.Bind(&renunganHarian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := POSTRenunganHarian(renunganHarian, foto)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTRenunganHarian(renunganHarian models.RenunganHarian, foto *multipart.FileHeader) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO renungan_harian (isi, foto) VALUES (?, ?)"

	result, err := con.Exec(sqlStatement, renunganHarian.Isi, renunganHarian.Foto)

	if err != nil {
		return res, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return res,err
	}

	folder := "renungan"

	newFotoName := fmt.Sprintf("%s-%d.png", folder, lastID)

	updateSQL := "UPDATE renungan_harian SET foto = ? WHERE id = ?"
	_, err = con.Exec(updateSQL, newFotoName, lastID)
	if err != nil {
		return res, err
	}

	_, uploadErr := UploadFotoFolder(foto, int64(lastID), folder)
	if uploadErr != nil {
		return res, uploadErr
	}

	res.Status = http.StatusCreated
	res.Message = "Renungan harian berhasil ditambahkan"
	res.Data = renunganHarian

	return res, nil
}

func SoftDeletedataRenunganHarian(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id"})
	}
	
	result, err := UpdateDeletedAtRenunganHarian(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtRenunganHarian(id int) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE renungan_harian SET deleted_at = NOW() WHERE id = ?"

	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Renungan harian berhasil dihapus"

	return res, nil
}

func EditRenunganHarian(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id"})
	}

	isi := c.FormValue("isi")
	foto, err := c.FormFile("foto")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	var filename string
	if err == nil{
		folder := "renungan"
		result, uploadErr := UploadFotoFolder(foto, int64(id), folder)
		if uploadErr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": uploadErr.Error()})
		}
		filename = result.Data.(map[string]string)["filename"]
	}


	renunganHarian := models.RenunganHarian{
		Id : id,
		Isi: isi,
		Foto: filename,
	}

	if err := c.Bind(&renunganHarian); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := PUTRenunganHarian(renunganHarian)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PUTRenunganHarian(renunganHarian models.RenunganHarian) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE renungan_harian SET isi = ?, foto = ? WHERE id = ?"

	_, err := con.Exec(sqlStatement, renunganHarian.Isi, renunganHarian.Foto, renunganHarian.Id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Renungan harian berhasil diupdate"

	return res, nil
}
