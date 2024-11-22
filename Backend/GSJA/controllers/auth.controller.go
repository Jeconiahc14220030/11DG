package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"database/sql"
	"errors"
	"fmt"
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

	sqlStatement := "SELECT id,username,password FROM anggota WHERE username = ?"
	row := con.QueryRow(sqlStatement, username)

	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, fmt.Errorf("user not found")
		}
		return res, err
	}

	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if err != nil {
	// 	fmt.Println("Password does not match:", err)
	// } else {
	// 	fmt.Println("Password matches!")
	// }

	if user.Password == password {
		fmt.Println("Password cocok")
		return res, nil
	} else {
		fmt.Println("Password tidak cocok")
		return res, errors.New("password tidak cocok")
	}
	

	res.Status = http.StatusOK
	res.Message = "Login successful"
	res.Data = user

	return res, nil
}