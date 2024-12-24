package controllers

import (
	"GSJA/db"
	"GSJA/models"
	_"database/sql"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllCarousel(c echo.Context) error {
	result, err := GETAllCarousel()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETAllCarousel() (models.Response, error) {
	var carousel models.Carousel
	var arrayCarousel []models.Carousel
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM carousel"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&carousel.Id, 
			&carousel.Foto,
			&carousel.Status, 
			&carousel.CreatedAt, 
			&carousel.UpdatedAt, 
			&carousel.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayCarousel = append(arrayCarousel, carousel)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayCarousel

	return response, err
}

func FetchCarouselById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := GETCarouselById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETCarouselById(id int) (models.Response, error) {
	var carousel models.Carousel
	var arrayCarousel []models.Carousel
	var response models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM carousel WHERE id = ?"

	rows, err := con.Query(sqlStatement, id)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&carousel.Id, 
			&carousel.Foto,
			&carousel.Status, 
			&carousel.CreatedAt, 
			&carousel.UpdatedAt, 
			&carousel.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayCarousel = append(arrayCarousel, carousel)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayCarousel

	return response, err
}

func AddCarousel(c echo.Context) error {
	foto, err := c.FormFile("foto")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	

	carousel := models.Carousel{
		Foto:  foto.Filename,
	}

	result, err := POSTCarousel(carousel, foto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTCarousel(carousel models.Carousel, foto *multipart.FileHeader) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO carousel (foto) VALUES (?)"
	result, err := con.Exec(sqlStatement, carousel.Foto)

	if err != nil {
		return res, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	folder := "carousel"
	
	newFotoName := fmt.Sprintf("%s-%d.png", folder, lastID)

	updateSQL := "UPDATE carousel SET foto = ? WHERE id = ?"
	_, err = con.Exec(updateSQL, newFotoName, lastID)
	if err != nil {
		return res, err
	}

	_, uploadErr := UploadFotoFolder(foto, int64(lastID), folder)
	if uploadErr != nil {
		return res, uploadErr
	}

	res.Status = http.StatusCreated
	res.Message = "Carousel added successfully"
	res.Data = carousel

	return res, nil
}

func UpdateDeletedatCarousel(id int) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	sqlStatement := "UPDATE carousel SET deleted_at = NOW() WHERE id = ?"
	_, err := con.Exec(sqlStatement, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Carousel deleted successfully"

	return res, nil
}

func SoftDeleteCarousel(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	result, err := UpdateDeletedatCarousel(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetActiveCaraousel() (models.Response, error) {
	var carousel models.Carousel
	var arrayCarousel []models.Carousel
	var response models.Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM carousel WHERE status = 'active'"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&carousel.Id, 
			&carousel.Foto,
			&carousel.Status, 
			&carousel.CreatedAt, 
			&carousel.UpdatedAt, 
			&carousel.DeletedAt,
		)

		if err != nil {
			return response, err
		}

		arrayCarousel = append(arrayCarousel, carousel)
	}

	response.Status = http.StatusOK
	response.Message = "OK"
	response.Data = arrayCarousel

	return response, err
}

func FetchActiveCarousel(c echo.Context) error {
	result, err := GetActiveCaraousel()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func EditStatusCarousel(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	status := c.FormValue("status")

	result, err := PUTStatusCarousel(id, status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PUTStatusCarousel(id int, status string) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()
	sqlStatement := "UPDATE carousel SET status = ? WHERE id = ?"
	_, err := con.Exec(sqlStatement, status, id)

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Carousel status updated successfully"

	return res, nil
}
