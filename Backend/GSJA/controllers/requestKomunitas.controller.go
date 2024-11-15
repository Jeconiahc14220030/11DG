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

func POSTRequestKomunitas(c echo.Context) error {
    var Requestkomunitas models.Requestkomunitas
    var res models.Response

    con := db.CreateCon()
    if err := c.Bind(&Requestkomunitas); err != nil {
        return err
    }

    sqlStatement := "INSERT INTO request_komunitas (id_komunitas, id_anggota) VALUES (?, ?)"

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }
    defer stmt.Close()

    result, err := stmt.Exec(Requestkomunitas.Id_komunitas, Requestkomunitas.Id_anggota)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    lastInsertId, err := result.LastInsertId()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    Requestkomunitas.Id = int(lastInsertId)

    res.Status = http.StatusOK
    res.Message = "Success"
    res.Data = Requestkomunitas

    return c.JSON(http.StatusOK, res)
}


func ReceiveRequestKomunitas(c echo.Context) error {
    var Requestkomunitas models.Requestkomunitas
    var res models.Response

    con := db.CreateCon()

    // Bind the incoming JSON data to the Requestkomunitas struct
    if err := c.Bind(&Requestkomunitas); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
    }

    // Ensure that the `id` and `status` fields are provided in the JSON
    if Requestkomunitas.Id == 0 || Requestkomunitas.Status == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Missing id or status"})
    }

    sqlStatement := "UPDATE request_komunitas SET status = ?, update_at = NOW() WHERE id_request = ?"

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }
    defer stmt.Close()

    // Execute the update query with the provided `status` and `id`
    _, err = stmt.Exec(Requestkomunitas.Status, Requestkomunitas.Id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    res.Status = http.StatusOK
    res.Message = "Status updated successfully"
    res.Data = Requestkomunitas

    return c.JSON(http.StatusOK, res)
}
