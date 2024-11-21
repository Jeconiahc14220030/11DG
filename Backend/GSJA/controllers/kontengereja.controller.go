package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchKontenGereja(c echo.Context) error {
	result, err := GETKontenGereja()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETKontenGereja() (models.Response, error) {
	var kontenGereja models.KontenGereja
	var arrayKontenGereja []models.KontenGereja
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM konten_gereja"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&kontenGereja.Id, 
			&kontenGereja.Visi, 
			&kontenGereja.Misi, 
			&kontenGereja.Tujuan, 
			&kontenGereja.PesanKetua, 
			&kontenGereja.IdCarousel, 
			&kontenGereja.IdBerita, 
			&kontenGereja.IdKutipanHarian, 
			&kontenGereja.IdRenunganHarian,
			&kontenGereja.CreatedAt,
			&kontenGereja.UpdatedAt,
			&kontenGereja.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayKontenGereja = append(arrayKontenGereja, kontenGereja)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayKontenGereja

	return response, err
}

func AddKontenGereja(c echo.Context) error {
	visi := c.FormValue("visi")
	misi := c.FormValue("misi")
	tujuan := c.FormValue("tujuan")
	pesanKetua := c.FormValue("pesanKetua")

	kontenGereja := models.KontenGereja{
		Visi: visi,
		Misi: misi,
		Tujuan: tujuan,
		PesanKetua: pesanKetua,
	}

	if err := c.Bind(&kontenGereja); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := POSTKontenGereja(kontenGereja)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTKontenGereja(kontenGereja models.KontenGereja) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT INTO konten_gereja (visi, misi, tujuan, pesan_ketua) VALUES (?, ?, ?, ?)"
	_, err := con.Exec(sqlStatement, kontenGereja.Visi, kontenGereja.Misi, kontenGereja.Tujuan, kontenGereja.PesanKetua)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Konten gereja added successfully"
	res.Data = kontenGereja

	return res, nil
}

func SoftDeletedataKontenGereja(c echo.Context) error {
	id := c.Param("id")

	result, err := UpdateDeletedAtKontenGereja(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDeletedAtKontenGereja(id string) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "UPDATE konten_gereja SET deleted_at = NOW() WHERE id = ?"
	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Konten gereja deleted successfully"

	return res, nil
}
