package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Username and password are required"})
	}

	result, err := AuthenticateUser (username, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	return c.JSON(http.StatusOK, result)
}

func AuthenticateUser (username, password string) (models.Response, error) {
	var user models.User
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "SELECT * FROM anggota WHERE username = ?"
	row := con.QueryRow(sqlStatement, username)

	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, fmt.Errorf("user not found")
		}
		return res, err
	}

	// Directly compare the stored password with the input password
	if user.Password != password {
		return res, fmt.Errorf("invalid password")
	}

	res.Status = http.StatusOK
	res.Message = "Login successful"
	res.Data = user

	return res, nil
}

func Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Username and password are required"})
	}

	user := models.User{Username: username, Password: password}

	// No hashing, store password as plain text
	result, err := POSTUser (user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTUser (user models.User) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT INTO anggota (username, password) VALUES (?, ?)"
	_, err := con.Exec(sqlStatement, user.Username, user.Password)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "User  registered successfully"
	res.Data = user

	return res, nil
}