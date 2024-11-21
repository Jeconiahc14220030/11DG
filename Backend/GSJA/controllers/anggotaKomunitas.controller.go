package controllers
import(
	"GSJA/db"
	"GSJA/models"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GETAllRequest() (models.Response, error) {
	var AnggotaKomunitas models.AnggotaKomunitas
	var arrayAnggotaKomunitas []models.AnggotaKomunitas
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM anggota_komunitas"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&AnggotaKomunitas.Id, 
			&AnggotaKomunitas.IdKomunitas, 
			&AnggotaKomunitas.IdAnggota,
			&AnggotaKomunitas.RequestAt,
			&AnggotaKomunitas.UpdatedAt,
			&AnggotaKomunitas.DeletedAt,
            &AnggotaKomunitas.Status, 
		)
		
		if err != nil {
			return res, err
		}

		arrayAnggotaKomunitas = append(arrayAnggotaKomunitas, AnggotaKomunitas)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrayAnggotaKomunitas

	return res, nil
}	

func FetchAllAnggotaKomunitas(c echo.Context) error {
	result, err := GETAllRequest()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func POSTAnggotaKomunitas(c echo.Context) error {
    var AnggotaKomunitas models.AnggotaKomunitas
    var res models.Response

    con := db.CreateCon()
    if err := c.Bind(&AnggotaKomunitas); err != nil {
        return err
    }

    sqlStatement := "INSERT INTO anggota_komunitas (id, id_anggota) VALUES (?, ?)"

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }
    defer stmt.Close()

    result, err := stmt.Exec(AnggotaKomunitas.Id, AnggotaKomunitas.Id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    lastInsertId, err := result.LastInsertId()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    AnggotaKomunitas.Id = int(lastInsertId)

    res.Status = http.StatusOK
    res.Message = "Success"
    res.Data = AnggotaKomunitas

    return c.JSON(http.StatusOK, res)
}


func ReceiveAnggotaKomunitas(c echo.Context) error {
    var AnggotaKomunitas models.AnggotaKomunitas
    var res models.Response

    con := db.CreateCon()

    // Bind the incoming JSON data to the AnggotaKomunitas struct
    if err := c.Bind(&AnggotaKomunitas); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
    }

    // Ensure that the `id` and `status` fields are provided in the JSON
    if AnggotaKomunitas.Id == 0 || AnggotaKomunitas.Status == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Missing id or status"})
    }

    sqlStatement := "UPDATE anggota_komunitas SET status = ?, update_at = NOW() WHERE id = ?"

    stmt, err := con.Prepare(sqlStatement)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }
    defer stmt.Close()

    // Execute the update query with the provided `status` and `id`
    _, err = stmt.Exec(AnggotaKomunitas.Status, AnggotaKomunitas.Id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
    }

    res.Status = http.StatusOK
    res.Message = "Status updated successfully"
    res.Data = AnggotaKomunitas

    return c.JSON(http.StatusOK, res)
}
