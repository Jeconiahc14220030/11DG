package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	// Struktur untuk menampung data JSON yang diterima

	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Println(username)
	fmt.Println(password)

	// Autentikasi pengguna
	result, err := AuthenticateUser(username, password)
	if err != nil {
		fmt.Println("error 3")
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	// Respon sukses
	return c.JSON(http.StatusOK, result)
}

func AuthenticateUser(username, password string) (models.Response, error) {
	var user models.User
	var res models.Response

	con := db.CreateCon()
	

	sqlStatement := "SELECT id, username, password FROM anggota WHERE username = ?"
	row := con.QueryRow(sqlStatement, username)

	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, fmt.Errorf("user not found")
		}
		return res, err
	}

	// defer con.Close()

	// if err != nil {
	// 	fmt.Println("Password does not match:", err)
	// } else {
	// 	fmt.Println("Password matches!")
	// }

	if user.Password == password {
		fmt.Println("Password cocok")
		res.Status = http.StatusOK
		res.Message = "Berhasil login"
		res.Data = user
		return res, nil
	} else {
		fmt.Println("Password tidak cocok")
		return res, errors.New("password tidak cocok")
	}
}