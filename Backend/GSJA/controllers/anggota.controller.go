package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllAnggota(c echo.Context) error {
	result, err := GETAllAnggota()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllAnggota() (models.Response, error) {
	var anggota models.Anggota
	var arrayAnggota []models.Anggota
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM anggota"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&anggota.Id, 
			&anggota.Nama, 
			&anggota.Username,
			&anggota.Password,
			&anggota.Email, 
			&anggota.NomorTelepon, 
			&anggota.TanggalLahir,
			&anggota.Poin,
			&anggota.CreatedAt,
			&anggota.UpdatedAt,
			&anggota.DeletedAt,
			&anggota.IdKomunitas,
			&anggota.IdHf,
		)
		
		if err != nil {
			return res, err
		}

		arrayAnggota = append(arrayAnggota, anggota)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayAnggota

	return res, nil
}

func FetchAnggotaById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GETAnggotaById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAnggotaById(id int) (models.Response, error) {
	var anggota models.Anggota
	var arrayAnggota []models.Anggota
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM anggota WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&anggota.Id, 
			&anggota.Nama, 
			&anggota.Email, 
			&anggota.NomorTelepon, 
			&anggota.TanggalLahir,
			&anggota.Poin,
			&anggota.CreatedAt,
			&anggota.UpdatedAt,
			&anggota.DeletedAt,
			&anggota.IdKomunitas,
			&anggota.IdHf,
		)

		if err != nil {
			return res, err
		}

		arrayAnggota = append(arrayAnggota, anggota)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayAnggota

	return res, nil
}

func FetchRiwayatVoucherByAnggotaId(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GETRiwayatVoucherAnggota(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETRiwayatVoucherAnggota(id int) (models.Response, error) {
	var riwayatVoucher models.HistoryPembelianVoucher
	var arrRiwayatVoucher []models.HistoryPembelianVoucher
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM history_pembelian_voucher WHERE id_anggota = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&riwayatVoucher.Id,
			&riwayatVoucher.Tanggal,
			&riwayatVoucher.CreatedAt,
			&riwayatVoucher.UpdatedAt,
			&riwayatVoucher.DeletedAt,
			&riwayatVoucher.IdAnggota,
			&riwayatVoucher.IdVoucher,
		)

		if err != nil {
			return response, err
		}

		arrRiwayatVoucher = append(arrRiwayatVoucher, riwayatVoucher)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrRiwayatVoucher

	return response, err
}

func AddAnggota(c echo.Context) error {
	var anggota models.Anggota

	if err := c.Bind(&anggota); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := POSTAnggota(anggota)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTAnggota(anggota models.Anggota) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT INTO anggota (nama, email, nomor_telepon, tanggal_lahir, poin, created_at, updated_at, id_komunitas, id_hf) VALUES (?, ?, ?, ?, ?, NOW(), NOW(), ?, ?)"
	_, err := con.Exec(sqlStatement, anggota.Nama, anggota.Email, anggota.NomorTelepon, anggota.TanggalLahir, anggota.Poin, anggota.IdKomunitas, anggota.IdHf)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Anggota added successfully"
	res.Data = anggota 

	return res, nil
}


