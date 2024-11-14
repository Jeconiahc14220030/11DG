package controllers

import (
	"GSJA/db"
	"GSJA/models"
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
