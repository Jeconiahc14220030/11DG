package routes

import (
	"GSJA/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins:     []string{"*"},                           // Allow all origins
        AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType}, // Allowed headers
        AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete}, // Allowed methods
        AllowCredentials: true,                                    // Allow credentials (cookies, authorization headers, etc.)
    }))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Berhasil terkoneksi dengan database!")
	})

	e.POST("/login", func(c echo.Context) error {
		return controllers.Login(c)
	})

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

	e.GET("/absensihf/anggota/:id", controllers.FetchAbsensihfByAnggotaId)

	e.GET("Anggotahf/:idHf", controllers.GetAnggotaHF)

	e.GET("/berita", controllers.FetchAllBerita)

	e.GET("/renunganharian", controllers.FetchAllRenunganHarian)

	e.GET("/kutipanharian", controllers.FetchAllKutipanHarian)

	e.GET("/anggota/:id/riwayatvoucher", controllers.FetchRiwayatVoucherByAnggotaId)

	e.GET("/hf", controllers.FetchAllHf)
	
	e.GET("/anggotakomunitas", controllers.FetchAllAnggotaKomunitas); e.GET("/anggotakomunitas/pending", controllers.FetchPendingRequest); e.GET("/anggotakomunitas/member", controllers.FetchAllMemberAnggotaKomunitas);

	e.GET("/jadwal/:id", controllers.FetchJadwalById)

	e.GET("/anggota/:id/absensi", controllers.FetchAbsensiById)

	e.GET(":username", controllers.FetchAnggotaByUsername)
	e.GET("/roles", controllers.FetchAllRoles)
	
	e.GET("anggotaRoles", controllers.FetchAllRolesAnggota)

	e.POST("/roles/add", controllers.AddRoles)

	e.POST("/anggota/add", controllers.AddAnggota)

	e.POST("renunganharian/add", controllers.AddRenunganHarian)

	e.POST("jadwal/add", controllers.AddJadwal) // perlu dipikirkan lagi format tanggal

	e.POST("/berita/add", controllers.AddBerita)

	e.POST("kutipanharian/add", controllers.AddKutipanHarian)

	e.POST("voucher/add", controllers.AddVoucher)

	e.POST("carousel/add", controllers.AddCarousel)

	e.PUT("anggotaKomunitas/updatestatus", controllers.UpdateRequestStatus)
 
	e.POST("laporankeuangan/add", controllers.AddLaporanKeuangan) // perlu dipikirkan lagi struktur tanggal 

	e.POST("absensihf/add", controllers.AddAbsensiHfForm) // perlu dipikirkan lagi struktur tanggal

	e.POST("jadwallatihan/add", controllers.AddJadwalLatihan) // perlu dipikirkan lagi struktur tanggal

	e.POST("komunitas/:id/pengumuman/tambah", controllers.AddPengumuman) // perlu dipikirkan lagi struktur tanggal, ga perlu tanggal ga si, tinggal ambil dari created_at

	e.POST("kontengereja/add", controllers.AddKontenGereja)

	e.PUT("anggota/delete/:id", controllers.SoftDeleteAnggota); e.PUT("anggota/restore/:id", controllers.RestoreDeletedAnggota);

	e.POST("anggotaKomunitas/request", controllers.RequestJoinKomunitas)

	e.POST("absensi/add", controllers.AddKehadiranAbsensi)

	e.PUT("anggota/:id/editprofil", controllers.EditProfil)

	e.POST("anggota/tukarvoucher", controllers.TukarVoucher)

	e.PUT("anggota/changePassword/:id", controllers.ChangePassword)
	return e
}
