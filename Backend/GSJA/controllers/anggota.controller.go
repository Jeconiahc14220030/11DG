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

	sqlStatement := "SELECT * FROM anggota WHERE deleted_at IS NULL"

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
	username := c.FormValue("username")
	password := c.FormValue("password")
	email := c.FormValue("email")
	nomorTelepon := c.FormValue("nomor_telepon")
	tanggalLahir := c.FormValue("tanggal_lahir")

	anggota := models.Anggota{
		Username:     username,
		Password:     password,
		Email:        email,
		NomorTelepon: nomorTelepon,
		TanggalLahir: tanggalLahir,
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

	sqlStatement := "INSERT INTO anggota (username, password, email, nomor_telepon, tanggal_lahir) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := con.Exec(sqlStatement, anggota.Username, anggota.Password, anggota.Email, anggota.NomorTelepon, anggota.TanggalLahir 	)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Anggota added successfully"
	res.Data = anggota 

	return res, nil
}

func SoftDeleteAnggota(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	result, err := UpdateDeletedAtAnggota(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtAnggota(id int) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE anggota SET deleted_at = NOW() WHERE id = ?"
	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Deleted at updated successfully"
	res.Data = map[string]int{"id": id}

	return res, nil
}

func RestoreDeletedAnggota(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	result, err := UpdateDeletedAtToNullAnggota(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtToNullAnggota(id int) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE anggota SET deleted_at = NULL WHERE id = ?"
	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Deleted at restored successfully"
	res.Data = map[string]int{"id": id}

	return res, nil
}

func FetchAbsensiById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam	)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	result, err := GetAbsensiById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAbsensiById(id int) (models.Response, error) {
	var absensi models.Absensi
	var arrAbsensi []models.Absensi
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM absensi WHERE id_anggota = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&absensi.Id,
			&absensi.Status,
			&absensi.IdAnggota,
			&absensi.IdJadwal,
			&absensi.CreatedAt,
			&absensi.UpdatedAt,
			&absensi.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrAbsensi = append(arrAbsensi, absensi)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrAbsensi

	return response, err
}

