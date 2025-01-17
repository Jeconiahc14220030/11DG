package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"database/sql"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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
			&anggota.FotoProfil,
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
			&anggota.FotoProfil,
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
	res.Data = arrayAnggota

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
			&anggota.FotoProfil,
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
	nama := c.FormValue("nama")
	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")
	tanggalLahir := c.FormValue("tanggal_lahir")
	nomorTelepon := c.FormValue("nomor_telepon")

	anggota := models.Anggota{
		Nama:         nama,
		Email:        email,
		Username:     username,
		Password:     password,
		TanggalLahir: tanggalLahir,
		NomorTelepon: nomorTelepon,
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

	sqlStatement := "INSERT INTO anggota (id_hf, nama, username, password, email, nomor_telepon, tanggal_lahir, foto_profile) VALUES (NULL, ?, ?, ?, ?, ?, ?, 'dummy.png')"
	_, err := con.Exec(sqlStatement, anggota.Nama, anggota.Username, anggota.Password, anggota.Email, anggota.NomorTelepon, anggota.TanggalLahir)

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

func TukarVoucher(c echo.Context) error {
	idAnggotaStr := c.FormValue("idAnggota")
	idVoucherStr := c.FormValue("idVoucher")

	log.Printf("Data diterima - idAnggota: %s, idVoucher: %s\n", idAnggotaStr, idVoucherStr)

	idAnggota, err := strconv.Atoi(idAnggotaStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	idVoucher, err := strconv.Atoi(idVoucherStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	voucher := models.HistoryPembelianVoucher {
		IdAnggota: idAnggota,
		IdVoucher: idVoucher,
	} 

	result, err := POSTPenukaranVoucher(voucher)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func POSTPenukaranVoucher(voucher models.HistoryPembelianVoucher) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO history_pembelian_voucher (tanggal, id_anggota, id_voucher) VALUES (NOW(), ?, ?)"
	_, err := con.Exec(sqlStatement, voucher.IdAnggota, voucher.IdVoucher)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Penukaran voucher berhasil ditambahkan"
	res.Data = voucher

	return res, nil
}

func EditProfil(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	nama := c.FormValue("nama")
	email := c.FormValue("email")
	tanggalLahir := c.FormValue("tanggal_lahir")
	nomorTelepon := c.FormValue("nomor_telepon")

	// Handle file upload for profile photo
	fotoFile, err := c.FormFile("photo")
	var filename string
	if err == nil {
		// Upload photo and get the filename
		folder := "profiles"
		result, uploadErr := UploadFotoFolder(fotoFile, int64(id), folder)
		if uploadErr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": uploadErr.Error()})
		}
		filename = result.Data.(map[string]string)["filename"]
	}

	// Update the profile in the database
	result, err := PUTProfilAnggota(id, nama, email, tanggalLahir, nomorTelepon, filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PUTProfilAnggota(id int, nama string, email string, tanggalLahir string, nomorTelepon string, fotoProfile string) (models.Response, error) {
	var res models.Response

	sqlStatement := "UPDATE anggota SET nama = ?, tanggal_lahir = ?, email = ?, nomor_telepon = ?, foto_profile = ?, updated_at = NOW() WHERE id = ?"
	fmt.Printf("Nama: %s, Email: %s, Tanggal Lahir: %s, Nomor Telepon: %s\n", nama, email, tanggalLahir, nomorTelepon)

	con := db.CreateCon()
	_, err := con.Exec(sqlStatement, nama, tanggalLahir, email, nomorTelepon, fotoProfile, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Profile updated successfully"
	res.Data = map[string]int{"id": id}

	return res, nil
}

func UploadFotoFolder(file *multipart.FileHeader, id int64, folder string) (models.Response, error) {
	var res models.Response
	log.Println("Upload Foto")

	src, err := file.Open()
	if err != nil {
		log.Println(err.Error())
		return res, err
	}
	defer src.Close()

	destinationFolder := "uploads/" + folder
	if _, err := os.Stat(destinationFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(destinationFolder, os.ModePerm); err != nil {
			log.Println(err.Error())
			return res, err
		}
	}

	filename := folder + "-" + strconv.Itoa(int(id)) + ".png"
	destinationPath := destinationFolder + "/" + filename

	dst, err := os.Create(destinationPath)
	if err != nil {
		log.Println(err.Error())
		return res, err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		log.Println(err.Error())
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses Upload Foto"
	res.Data = map[string]string{"filename": filename}

	return res, nil
}
