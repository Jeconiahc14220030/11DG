package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"strconv"
	"time"

	// "strconv"
	"github.com/labstack/echo/v4"
)

func FetchAllJadwalLatihan(c echo.Context) error {
	result, err := GETJadwalLatihan()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETJadwalLatihan() (models.Response, error) {
	var jadwalLatihan models.JadwalLatihan
	var arrJadwalLatihan []models.JadwalLatihan
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM jadwal_latihan"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&jadwalLatihan.Id,
			&jadwalLatihan.Tanggal,
			&jadwalLatihan.Lokasi,
			&jadwalLatihan.CreatedAt,
			&jadwalLatihan.UpdatedAt,
			&jadwalLatihan.DeletedAt,
			&jadwalLatihan.IdAnggota,
			&jadwalLatihan.IdKomunitas,
		)

		if err != nil {
			return response, err
		}

		arrJadwalLatihan = append(arrJadwalLatihan, jadwalLatihan)
	}

	response.Status = http.StatusOK
	response.Message = "Berhasil GET semua jadwal latihan"
	response.Data = arrJadwalLatihan

	return response, err
}

func AddJadwalLatihan(c echo.Context) error {
    // Ambil data dari form
    tanggalStr := c.FormValue("tanggal") // Tanggal dari form dalam string
    lokasi := c.FormValue("lokasi")
    idAnggotaStr := c.FormValue("id_anggota")
    idKomunitasStr := c.FormValue("id_komunitas")

    // Parse idAnggota dan idKomunitas ke integer
    idAnggota, err := strconv.Atoi(idAnggotaStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id_anggota"})
    }

    idKomunitas, err := strconv.Atoi(idKomunitasStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id_komunitas"})
    }

    // Parse tanggal ke format YYYY-MM-DD
    tanggal, err := time.Parse("2006-01-02", tanggalStr) // "2006-01-02" adalah layout Go untuk YYYY-MM-DD
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid date format. Use YYYY-MM-DD"})
    }

    // Buat instance dari JadwalLatihan
    jadwalLatihan := models.JadwalLatihan{
        Tanggal:    tanggal,
        Lokasi:     lokasi,
        IdAnggota:  idAnggota,
        IdKomunitas: idKomunitas,
    }

    // Simpan ke database
    result, err := POSTJadwalLatihan(jadwalLatihan)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    return c.JSON(http.StatusCreated, result)
}

	

func POSTJadwalLatihan(jadwalLatihan models.JadwalLatihan) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO jadwal_latihan (tanggal, lokasi, id_anggota, id_komunitas) VALUES (?, ?, ?, ?)"
	_, err := con.Exec(sqlStatement, jadwalLatihan.Tanggal, jadwalLatihan.Lokasi, jadwalLatihan.IdAnggota, jadwalLatihan.IdKomunitas)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Jadwal latihan added successfully"
	res.Data = jadwalLatihan

	return res, nil
}

func SoftDeletedataJadwalLatihan(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id"})
	}

	result, err := UpdateDeletedAtJadwalLatihan(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtJadwalLatihan(id int) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "UPDATE jadwal_latihan SET deleted_at = NOW() WHERE id = ?"

	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Jadwal latihan berhasil dihapus"

	return res, nil
}