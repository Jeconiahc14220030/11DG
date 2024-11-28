package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllKomunitas(c echo.Context) error {
	result, err := GETAllKomunitas()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllKomunitas() (models.Response, error) {
	var komunitas models.KomunitasPelayanan
	var arrayKomunitas []models.KomunitasPelayanan
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM komunitas_pelayanan"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&komunitas.Id,
			&komunitas.NamaKomunitas,
			&komunitas.Deskripsi,
			&komunitas.Snk,
			&komunitas.Manfaat,
			&komunitas.Gambar,
			&komunitas.CreatedAt,
			&komunitas.UpdatedAt,
			&komunitas.DeletedAt,
		)

		if err != nil {
			return response, nil
		}

		arrayKomunitas = append(arrayKomunitas, komunitas)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayKomunitas

	return response, err
}

func FetchKomunitasById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GETKomunitasById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETKomunitasById(id int) (models.Response, error) {
	var komunitas models.KomunitasPelayanan
	var arrayKomunitas []models.KomunitasPelayanan
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM komunitas_pelayanan WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&komunitas.Id,
			&komunitas.NamaKomunitas,
			&komunitas.Deskripsi,
			&komunitas.Snk,
			&komunitas.Manfaat,
			&komunitas.Gambar,
			&komunitas.CreatedAt,
			&komunitas.UpdatedAt,
			&komunitas.DeletedAt,
		)

		if err != nil {
			return response, nil
		}

		arrayKomunitas = append(arrayKomunitas, komunitas)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayKomunitas

	return response, err
}

func AddPengumuman(c echo.Context) error {
	// konten := c.FormValue("konten")
	// tanggal := c.FormValue("tanggal")
	// strIdKomunitas := c.FormValue("id_komunitas")

	// idKomunitas, err := strconv.Atoi(strIdKomunitas)

	// formattanggal, err := time.Parse("2024-11-20", tanggal)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Tanggal"})
	// }

	// pengumuman := models.Pengumuman{
	// 	Konten: konten,
	// 	// Tanggal: tanggal,
	// 	Id_komunitas: idKomunitas,
	// }

	// idParam := c.Param("id")
	// id, err := strconv.Atoi(idParam)

	// if err := c.Bind(&pengumuman); err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	// }

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	// }

	// result, err := POSTPengumuman(pengumuman, id)

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	// }

	// return c.JSON(http.StatusCreated, result)
	return c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}

func POSTPengumuman(pengumuman models.Pengumuman, id int) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT INTO pengumuman(konten, tanggal, id_komunitas) VALUES (?, ?, ?)"
	_, err := con.Exec(sqlStatement, pengumuman.Konten, pengumuman.Tanggal, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Pengumuman berhasil ditambahkan"
	res.Data = pengumuman

	return res, nil
}
