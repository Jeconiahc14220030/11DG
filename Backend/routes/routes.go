package routes

import (
	"GSJA/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Berhasil terkoneksi dengan database!")
	})

	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)

	// Fetch Data
	e.GET("/anggota", controllers.FetchAllAnggota)
	e.GET("/anggota/:id", controllers.FetchAnggotaById)

	e.GET("/jadwal", controllers.FetchAllJadwal)

	e.GET("/jadwallatihan", controllers.FetchAllJadwalLatihan)

	e.GET("/komunitas", controllers.FetchAllKomunitas)
	e.GET("/komunitas/:id", controllers.FetchKomunitasById)

	e.GET("/voucher", controllers.FetchAllVoucher)
	e.GET("/voucher/:id", controllers.FetchVoucherById)

	e.GET("/dashboard", controllers.FetchKontenGereja)

	e.GET("/laporankeuangan", controllers.FetchAllLaporanKeuangan)
	e.GET("/laporankeuangan/:id", controllers.FetchLaporanKeuanganById)

	e.GET("/carousel", controllers.FetchAllCarousel)
	e.GET("/carousel/:id", controllers.FetchCarouselById)

	e.GET("/absensihf", controllers.FetchAllAbsensihf)
	e.GET("/absensihf/:id", controllers.FetchAbsensihfById)

	e.GET("/berita", controllers.FetchAllBerita)

	e.GET("/renunganharian", controllers.FetchAllRenunganHarian)

	e.GET("/kutipanharian", controllers.FetchAllKutipanHarian)

	e.GET("/anggota/:id/riwayatvoucher", controllers.FetchRiwayatVoucherByAnggotaId)

	e.POST("/anggota/add", controllers.AddAnggota)

	return e
}
