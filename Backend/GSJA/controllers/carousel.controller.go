package controllers

import (
	"GSJA/db"
	"GSJA/models"
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
			&carousel.Foto1, 
			&carousel.Foto2, 
			&carousel.Foto3, 
			&carousel.Foto4, 
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
			&carousel.Foto1, 
			&carousel.Foto2, 
			&carousel.Foto3, 
			&carousel.Foto4, 
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
	foto1 := c.FormValue("foto1")
	foto2 := c.FormValue("foto2")
	foto3 := c.FormValue("foto3")
	foto4 := c.FormValue("foto4")

	carousel := models.Carousel{
		Foto1:        foto1,
		Foto2:        foto2,
		Foto3:        foto3,
		Foto4:        foto4,
	}

	result, err := POSTCarousel(carousel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func POSTCarousel(carousel models.Carousel) (models.Response, error) {
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO carousel (foto1, foto2, foto3, foto4, status_carousel) VALUES (?, ?, ?, ?)"
	_, err := con.Exec(sqlStatement, carousel.Foto1, carousel.Foto2, carousel.Foto3, carousel.Foto4)

	if err != nil {
		return res, err
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
