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

func GetPendingRequest() (models.Response, error) {
	var AnggotaKomunitas models.AnggotaKomunitas
	var arrayAnggotaKomunitas []models.AnggotaKomunitas
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM anggota_komunitas WHERE status = 'pending'"

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

func FetchPendingRequest(c echo.Context) error {
	result, err := GetPendingRequest()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GETMember() (models.Response, error) {
	var AnggotaKomunitas models.AnggotaKomunitas
	var arrayAnggotaKomunitas []models.AnggotaKomunitas
	var res models.Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM anggota_komunitas WHERE status = 'diterima'"

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

func FetchAllMemberAnggotaKomunitas(c echo.Context) error {
	result, err := GETMember()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func RequestJoinKomunitas(c echo.Context) error {
    // Bind data JSON dari request body ke struct AnggotaKomunitas
    anggotaKomunitas := models.AnggotaKomunitas{}
    if err := c.Bind(&anggotaKomunitas); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
    }

    // Validasi input
    if anggotaKomunitas.IdAnggota == 0 || anggotaKomunitas.IdKomunitas == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "id_anggota and id_komunitas are required"})
    }

    // Database connection
    con := db.CreateCon()

    // Cek apakah pasangan id_anggota dan id_komunitas sudah ada
    sqlStatement := "SELECT COUNT(*) FROM anggota_komunitas WHERE id_anggota = ? AND id_komunitas = ?"
    var count int
    err := con.QueryRow(sqlStatement, anggotaKomunitas.IdAnggota, anggotaKomunitas.IdKomunitas).Scan(&count)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to check existing data"})
    }

    if count > 0 {
        return c.JSON(http.StatusConflict, map[string]string{"message": "Data already exists"})
    }

    // Insert data baru ke tabel anggota_komunitas
    sqlStatement = "INSERT INTO anggota_komunitas (id_anggota, id_komunitas, status) VALUES (?, ?, ?)"
    _, err = con.Exec(sqlStatement, anggotaKomunitas.IdAnggota, anggotaKomunitas.IdKomunitas, "pending")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create request"})
    }

    // Respon sukses
    return c.JSON(http.StatusOK, map[string]string{"message": "Request to join komunitas created successfully"})
}

func UpdateRequestStatus(c echo.Context) error {
	// Bind data JSON dari request body ke struct AnggotaKomunitas
	anggotaKomunitas := models.AnggotaKomunitas{}
	if err := c.Bind(&anggotaKomunitas); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Database connection
	con := db.CreateCon()

	// Update status request
	sqlStatement := "UPDATE anggota_komunitas SET status = ? WHERE id = ?"
	_, err := con.Exec(sqlStatement, anggotaKomunitas.Status, anggotaKomunitas.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update status"})
	}

	// Respon sukses
	return c.JSON(http.StatusOK, map[string]string{"message": "Request status updated successfully"})
}
