package controllers

import (
	"GSJA/db"
	"GSJA/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Login handles user login
func Login(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := AuthenticateUser(user.Username, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	return c.JSON(http.StatusOK, result)
}

// AuthenticateUser  checks the username and password
func AuthenticateUser (username, password string) (models.Response, error) {
	var user models.User
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "SELECT * FROM users WHERE username = ?"
	row := con.QueryRow(sqlStatement, username)

	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return res, err // User not found
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return res, err // Password does not match
	}

	res.Status = http.StatusOK
	res.Message = "Login successful"
	res.Data = user

	return res, nil
}

// Register handles user registration
func Register(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// Hash the password before saving it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not hash password"})
	}
	user.Password = string(hashedPassword)

	result, err := POSTUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

// POSTUser  saves a new user to the database
func POSTUser (user models.User) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := con.Exec(sqlStatement, user.Username, user.Password)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "User  registered successfully"
	res.Data = user

	return res, nil
}