package router

import (
	hanamnesis "rme/internal/handler/anamnesis"
	hantrian "rme/internal/handler/antrian"
	hauth "rme/internal/handler/auth"
	hdiag "rme/internal/handler/diagnosis"
	hdoctor "rme/internal/handler/doctor"
	hicd "rme/internal/handler/icd"
	hlok "rme/internal/handler/lokalis_bedah"
	hpatient "rme/internal/handler/patient"
	hpe "rme/internal/handler/pemeriksaan_ekg"
	hpfo "rme/internal/handler/pemeriksaan_fungsi_organ"
	hpl "rme/internal/handler/pemeriksaan_laboratorium"
	hppb "rme/internal/handler/pemeriksaan_penunjang_bedah"
	hpv "rme/internal/handler/pemeriksaan_vital"
	hpoli "rme/internal/handler/poli"
	hrencana "rme/internal/handler/rencana_rawat"
	hru "rme/internal/handler/rujuk_ulang"
	hta "rme/internal/handler/tatalaksana"

	"github.com/gin-gonic/gin"
)

// tatalaksana handler import

// RegisterAntrianRoutes mendaftarkan endpoint /antrian ke Gin router.
func RegisterAntrianRoutes(r *gin.Engine, handler *hantrian.Handler) {
	group := r.Group("/antrian")
	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.PATCH("/:id", handler.Update)
	group.PATCH("/patient/:id_pasien", handler.UpdateByPatient)
	group.DELETE("/patient/:id_pasien", handler.DeleteByPatient)
}

// RegisterAnamnesisRoutes mendaftarkan endpoint /anamnesis ke Gin router.
// Contoh pemakaian di main:
//
//	r := gin.Default()
//	handler := anamnesis.NewHandler(uc)
//	router.RegisterAnamnesisRoutes(r, handler)
//	r.Run(":8081")

func RegisterAnamnesisRoutes(r *gin.Engine, handler *hanamnesis.Handler) {
	group := r.Group("/anamnesis")

	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.GET("/:id_pasien", handler.GetByID)

	// Soft-hide endpoint
	group.PATCH("/:id/hide", handler.Hide)
	// Generic update (partial)
	group.PATCH("/:id", handler.Update)
}

// RegisterAuthRoutes mendaftarkan endpoint /auth ke Gin router.
func RegisterAuthRoutes(r *gin.Engine, handler *hauth.Handler) {
	group := r.Group("/auth")

	group.POST("/login", handler.Login)
}

// RegisterPatientRoutes mendaftarkan endpoint /patients ke Gin router.
func RegisterPatientRoutes(r *gin.Engine, handler *hpatient.Handler) {
	group := r.Group("/patients")

	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
}

// RegisterDoctorRoutes mendaftarkan endpoint /dokter ke Gin router.
func RegisterDoctorRoutes(r *gin.Engine, handler *hdoctor.Handler) {
	group := r.Group("/dokter")

	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)

	// Kompatibilitas endpoint lama
	groupLegacy := r.Group("/doctors")
	groupLegacy.POST("/", handler.Create)
	groupLegacy.GET("/", handler.GetAll)
}

// RegisterPoliRoutes mendaftarkan endpoint /poli ke Gin router.
func RegisterPoliRoutes(r *gin.Engine, handler *hpoli.Handler) {
	group := r.Group("/poli")
	group.GET("/", handler.GetAll)
}

// RegisterPemeriksaanVitalRoutes mendaftarkan endpoint /pemeriksaan_vital
func RegisterPemeriksaanVitalRoutes(r *gin.Engine, handler *hpv.Handler) {
	group := r.Group("/pemeriksaan_vital")
	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	// Generic update (partial)
	group.PATCH("/:id", handler.Update)
	// Soft-hide endpoint
	group.PATCH("/:id/hide", handler.Hide)
}

// RegisterPemeriksaanLaboratoriumRoutes mendaftarkan endpoint /pemeriksaan_laboratorium
func RegisterPemeriksaanLaboratoriumRoutes(r *gin.Engine, handler *hpl.Handler) {
	group := r.Group("/pemeriksaan_laboratorium")
	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.PATCH("/:id", handler.Update)
	group.DELETE("/:id", handler.Delete)
	group.PATCH("/:id/hide", handler.Hide)
}

// RegisterPemeriksaanEkgRoutes mendaftarkan endpoint /pemeriksaan_ekg
func RegisterPemeriksaanEkgRoutes(r *gin.Engine, handler *hpe.Handler) {
	group := r.Group("/pemeriksaan_ekg")
	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.PATCH("/:id", handler.Update)
	group.PATCH("/:id/hide", handler.Hide)
	group.DELETE("/:id", handler.Delete)
}

// RegisterICDRoutes mendaftarkan endpoint /icd10
func RegisterICDRoutes(r *gin.Engine, handler *hicd.Handler) {
	group := r.Group("/icd10")
	group.GET("/", handler.GetAll)
}

// RegisterTatalaksanaRoutes mendaftarkan endpoint /tatalaksana
func RegisterTatalaksanaRoutes(r *gin.Engine, handler *hta.Handler) {
	group := r.Group("/tatalaksana")
	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.PATCH("/:id", handler.Update)
	group.PATCH("/:id/hide", handler.Hide)
}

// RegisterLokalisBedahRoutes mendaftarkan endpoint /lokalis_bedah
func RegisterLokalisBedahRoutes(r *gin.Engine, handler *hlok.Handler) {
	group := r.Group("/lokalis_bedah")
	group.POST("/", handler.Create)
	group.GET("/:id", handler.GetByID)
	group.GET("/", handler.GetAll)
	group.PATCH(":id", handler.Update)
	group.PATCH("/:id/hide", handler.Hide)
}

// RegisterDiagnosisRoutes mendaftarkan endpoint /diagnosis
func RegisterDiagnosisRoutes(r *gin.Engine, handler *hdiag.Handler) {
	group := r.Group("/diagnosis")
	group.POST("/", handler.Create)
	group.GET("/:id", handler.GetByID)
	group.GET("/", handler.GetAll)
	group.PATCH(":id", handler.Update)
	// Soft-hide endpoint
	group.PATCH("/:id/hide", handler.Hide)
}

// RegisterPemeriksaanFungsiOrganRoutes mendaftarkan endpoint /pemeriksaan_fungsi_organ
func RegisterPemeriksaanFungsiOrganRoutes(r *gin.Engine, handler *hpfo.Handler) {
	group := r.Group("/pemeriksaan_fungsi_organ")
	group.POST("/", handler.Create)
	group.GET("/:id", handler.GetByID)
	group.GET("/", handler.GetAll)
	group.PATCH("/:id", handler.Update)
	group.PATCH("/:id/hide", handler.Hide)
}

// RegisterPemeriksaanPenunjangBedahRoutes mendaftarkan endpoint /pemeriksaan_penunjang_bedah
func RegisterPemeriksaanPenunjangBedahRoutes(r *gin.Engine, handler *hppb.Handler) {
	group := r.Group("/pemeriksaan_penunjang_bedah")
	group.POST("/", handler.Create)
	group.GET("/:id", handler.GetByID)
	group.GET("/", handler.GetAll)
	group.PATCH("/:id", handler.Update)
	group.PATCH("/:id/hide", handler.Hide)
	group.PATCH("/patient/:id_pasien/hide", handler.HideByPatient)
}

// RegisterRujukUlangRoutes mendaftarkan endpoint /rujuk_ulang
func RegisterRujukUlangRoutes(r *gin.Engine, handler *hru.Handler) {
	group := r.Group("/rujuk_ulang")
	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.PATCH("/:id", handler.Update)
	group.DELETE("/:id", handler.Delete)
	group.PATCH("/:id/hide", handler.Hide)
}

// RegisterRencanaRawatRoutes mendaftarkan endpoint /rencana_rawat
func RegisterRencanaRawatRoutes(r *gin.Engine, handler *hrencana.Handler) {
	group := r.Group("/rencana_rawat")
	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.GET("/:id", handler.GetByID)
	group.PATCH("/:id", handler.Update)
	group.PATCH("/:id/hide", handler.Hide)
}
