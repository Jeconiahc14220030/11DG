package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)


func GetAllRoles() (models.Response, error) {
	var roles models.Roles
	var arrayRoles []models.Roles
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM roles"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&roles.Id, 
			&roles.Roles, 
			&roles.CreatedAt, 
			&roles.UpdatedAt, 
			&roles.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayRoles = append(arrayRoles, roles)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayRoles

	return response, nil
}

func FetchAllRoles(c echo.Context) error {
	result, err := GetAllRoles()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}


func AddRoles(c echo.Context) error {
	rolesName := c.FormValue("roles")

	roles := models.Roles{
		Roles: rolesName,
	}

	result, err := POSTRoles(roles)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func POSTRoles(roles models.Roles) (models.Response, error) {
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO roles (roles) VALUES (?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return response, err
	}

	result, err := stmt.Exec(roles.Roles)

	if err != nil {
		return response, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return response, err
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = map[string]int64{
		"last_insert_id": lastInsertId,
	}

	return response, nil
}
