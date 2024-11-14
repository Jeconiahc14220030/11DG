package controllers
import(
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GETAllRequest() (models.Response, error) {
	var Requestkomunitas models.Requestkomunitas
	var arrayRequestkomunitas []models.Requestkomunitas
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM request_komunitas"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&Requestkomunitas.Id, 
			&Requestkomunitas.Id_komunitas, 
			&Requestkomunitas.Id_anggota,
			&Requestkomunitas.RequestAt,
			&Requestkomunitas.Status, 
			&Requestkomunitas.UpdatedAt,
			&Requestkomunitas.DeletedAt,
		)
		
		if err != nil {
			return res, err
		}

		arrayRequestkomunitas = append(arrayRequestkomunitas, Requestkomunitas)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayRequestkomunitas

	return res, nil
}	

func FetchAllRequestKomunitas(c echo.Context) error {
	result, err := GETAllRequest()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}