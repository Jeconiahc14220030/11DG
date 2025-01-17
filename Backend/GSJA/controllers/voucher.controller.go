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

func FetchAllVoucher(c echo.Context) error {
	result, err := GETAllVoucher()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllVoucher() (models.Response, error) {
	var voucher models.Voucher
	var arrayVoucher []models.Voucher
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM voucher"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&voucher.Id, 
			&voucher.NamaVoucher, 
			&voucher.Status, 
			&voucher.Harga, 
			&voucher.Foto, 
			&voucher.CreatedAt, 
			&voucher.UpdatedAt, 
			&voucher.DeletedAt,
		)

		if err != nil {
			return response, err
		}

	arrayVoucher = append(arrayVoucher, voucher)
}

response.Status = http.StatusOK
response.Message = "OK"
response.Data = arrayVoucher

return response, err
}

func FetchVoucherById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GETVoucherById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func AddVoucher(c echo.Context) error {
	namaVoucher := c.FormValue("nama_voucher")
	status := c.FormValue("status")
	hargaStr := c.FormValue("harga")
	harga, err := strconv.Atoi(hargaStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Harga harus berupa angka"})
	}
	foto,err := c.FormFile("foto")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	fotoName := foto.Filename
	
	result, err := InsertVoucher(namaVoucher, status, harga, fotoName, foto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func InsertVoucher(namaVoucher, status string, harga int, fotoName string, foto *multipart.FileHeader) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := `
		INSERT INTO voucher (nama_voucher, status, harga, foto, created_at, updated_at)
		VALUES (?, ?, ?, ?, NOW(), NOW())
	`
	result, err := con.Exec(sqlStatement, namaVoucher, status, harga, fotoName)
	if err != nil {
		return res, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	
	folder := "voucher"

	newFotoName := fmt.Sprintf("%s-%d.png", folder, lastID)

	updateSQL := `
		UPDATE voucher
		SET foto = ?
		WHERE id = ?
	`
	_, err = con.Exec(updateSQL, newFotoName, lastID)
	if err != nil {
		return res, err
	}

	_, uploadErr := UploadFotoFolder(foto, int64(lastID), folder)
	if uploadErr != nil {
		return res, uploadErr
	}


	res.Status = http.StatusOK
	res.Message = "Voucher berhasil ditambahkan"
	return res, nil
}



func GETVoucherById(id int) (models.Response, error) {
	var voucher models.Voucher
	var arrayVoucher []models.Voucher
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM voucher WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&voucher.Id, 
			&voucher.NamaVoucher, 
			&voucher.Status, 
			&voucher.Harga, 
			&voucher.Foto, 
			&voucher.CreatedAt, 
			&voucher.UpdatedAt, 
			&voucher.DeletedAt,
		)

		if err != nil {
			return response, err
		}



		arrayVoucher = append(arrayVoucher, voucher)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayVoucher

	return response, err
}

func SoftDeleteVoucher(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := UpdateDeletedAtVoucher(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtVoucher(id int) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE voucher SET deleted_at = NOW() WHERE id = ?"
	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Voucher deleted successfully"

	return res, nil
}

func EditVoucher(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	namaVoucher := c.FormValue("nama_voucher")
	status := c.FormValue("status")
	hargaStr := c.FormValue("harga")
	harga, err := strconv.Atoi(hargaStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Harga harus berupa angka"})
	}
	foto,err := c.FormFile("foto")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	var filename string
	if err == nil{
		folder := "voucher"
		result, uploadErr := UploadFotoFolder(foto, int64(id), folder)
		if uploadErr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": uploadErr.Error()})
		}
		filename = result.Data.(map[string]string)["filename"]
	}

	result, err := PUTVoucher(id, namaVoucher, status, harga, filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PUTVoucher(id int, namaVoucher, status string, harga int, foto string) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := `
		UPDATE voucher
		SET nama_voucher = ?, status = ?, harga = ?, foto = ?, updated_at = NOW()
		WHERE id = ?
	`
	_, err := con.Exec(sqlStatement, namaVoucher, status, harga, foto, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Voucher updated successfully"
	return res, nil
}