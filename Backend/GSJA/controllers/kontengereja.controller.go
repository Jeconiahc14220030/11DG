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
