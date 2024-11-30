package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
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
	// tanggal := c.FormValue("tanggal")
	// lokasi := c.FormValue("lokasi")
	// strIdAnggota := c.FormValue("id_anggota")
	// strIdKomunitas := c.FormValue("id_komunitas")

	// // idAnggota, err := strconv.Atoi(strIdAnggota)

	// // if err != nil {
	// // 	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid id anggota"})
	// // }

	// // idKomunitas, err := strconv.Atoi(strIdKomunitas)

	// formattanggal, err := time.Parse("2024-11-20", tanggal)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Tanggal"})
	// }

	// idKomunitas, err := strconv.Atoi(strIdKomunitas)

	// formattanggal, err := time.Parse("2024-11-20", tanggal)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Tanggal"})
	// }

	// jadwalLatihan := models.JadwalLatihan{
	// 	// Tanggal: tanggal,
	// 	Lokasi: lokasi,
	// 	IdAnggota: idAnggota,
	// 	IdKomunitas: idKomunitas,
	// }

	// if err := c.Bind(&jadwalLatihan); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	// }

	// result, err := POSTJadwalLatihan(jadwalLatihan)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	// }

	// return c.JSON(http.StatusCreated, result)
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}

func POSTJadwalLatihan(jadwalLatihan models.JadwalLatihan) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

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
	id := c.Param("id")

	result, err := UpdateDeletedAtJadwalLatihan(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtJadwalLatihan(id string) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "UPDATE jadwal_latihan SET deleted_at = NOW() WHERE id = ?"

	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Jadwal latihan berhasil dihapus"

	return res, nil
}