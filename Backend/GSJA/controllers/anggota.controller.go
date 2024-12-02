package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"database/sql"
	"fmt"
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
			&anggota.IdHf,
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
			&anggota.IdHf,
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
		)

		if err != nil {
			return res, err
		}
		res.Data = anggota
		arrayAnggota = append(arrayAnggota, anggota)
	}

	res.Status = http.StatusOK
	res.Message = "Success"

	return res, nil
}

func FetchAnggotaByUsername(c echo.Context) error {
	username := c.Param("username")

	result, err := GETAnggotaByUsername(username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAnggotaByUsername(username string) (models.Response, error) {
	var anggota models.Anggota
	var arrayAnggota []models.Anggota
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM anggota WHERE username = ?"

	rows, err := con.Query(sqlStatement, username)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&anggota.Id, 
			&anggota.IdHf,
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
		)

		if err != nil {
			return res, err
		}
		res.Data = anggota
		arrayAnggota = append(arrayAnggota, anggota)
	}

	res.Status = http.StatusOK
	res.Message = "Success"

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
	nama := c.FormValue("nama")

	anggota := models.Anggota{
		Username:     username,
		Password:     password,
		Email:        email,
		NomorTelepon: nomorTelepon,
		TanggalLahir: tanggalLahir,
		Nama:         nama,
	}

	result, err := POSTAnggota(anggota)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTAnggota(anggota models.Anggota) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO anggota (username, nama, password, email, nomor_telepon, tanggal_lahir) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())"
	_, err := con.Exec(sqlStatement,anggota.Nama, anggota.Username, anggota.Password, anggota.Email, anggota.NomorTelepon, anggota.TanggalLahir)

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
	id, err := strconv.Atoi(idParam)

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

func ChangePassword(c echo.Context) error {
	idParam := c.Param("id")
	currentPassword := c.FormValue("current_password")
	newPassword := c.FormValue("new_password")
	confirmPassword := c.FormValue("confirm_password")

	
	if newPassword != confirmPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "New password and confirm password do not match"})
	}

	
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID"})
	}

	isValid, err := ValidateCurrentPassword(id, currentPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if !isValid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Current password is incorrect"})
	}

	result, err := UpdatePassword(id, newPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ValidateCurrentPassword(id int, currentPassword string) (bool, error) {
	var passwordFromDB string

	con := db.CreateCon()
	sqlStatement := "SELECT password FROM anggota WHERE id = ? AND deleted_at IS NULL"

	err := con.QueryRow(sqlStatement, id).Scan(&passwordFromDB)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil 
		}
		return false, err
	}

	if currentPassword != passwordFromDB {
		return false, nil
	}

	return true, nil
}

func UpdatePassword(id int, newPassword string) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	sqlStatement := "UPDATE anggota SET password = ?, updated_at = NOW() WHERE id = ?"

	_, err := con.Exec(sqlStatement, newPassword, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Password updated successfully"
	res.Data = map[string]int{"id": id}

	return res, nil
}

func EditProfil(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	nama := c.FormValue("nama")
	tanggalLahir := c.FormValue("tanggal_lahir")
	email := c.FormValue("email")
	nomorTelepon := c.FormValue("nomor_telepon")

	result, err := PUTProfilAnggota(id, nama, tanggalLahir, email, nomorTelepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PUTProfilAnggota(id int, nama string, tanggalLahir string, email string, nomorTelepon string) (models.Response, error) {
	var res models.Response

	sqlStatement := "UPDATE anggota SET nama = ?, tanggal_lahir = ?, email = ?, nomor_telepon = ?, updated_at = NOW() WHERE id = ?"

	con := db.CreateCon()
	_, err := con.Exec(sqlStatement, nama, tanggalLahir, email, nomorTelepon, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Profile updated successfully"
	res.Data = map[string]int{"id": id}

	return res, nil
}