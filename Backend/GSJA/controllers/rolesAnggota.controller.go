package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllRolesAnggota() (models.Response, error) {
	var rolesAnggota models.RolesAnggota
	var arrayRolesAnggota []models.RolesAnggota
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM roles_anggota"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&rolesAnggota.Id,
			&rolesAnggota.IdAnggota,
			&rolesAnggota.IdRoles,
			&rolesAnggota.CreatedAt,
			&rolesAnggota.UpdatedAt,
			&rolesAnggota.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayRolesAnggota = append(arrayRolesAnggota, rolesAnggota)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayRolesAnggota

	return response, nil
}

func FetchAllRolesAnggota(c echo.Context) error {
	result, err := GetAllRolesAnggota()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}


