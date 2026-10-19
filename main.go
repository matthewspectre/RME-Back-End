package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	handlerAnamnesis "rme/internal/handler/anamnesis"
	handlerAntrian "rme/internal/handler/antrian"
	handlerAuth "rme/internal/handler/auth"
	handlerDiagnosis "rme/internal/handler/diagnosis"
	handlerDoctor "rme/internal/handler/doctor"
	handlerICD "rme/internal/handler/icd"
	handlerLokalisBedah "rme/internal/handler/lokalis_bedah"
	handlerPatient "rme/internal/handler/patient"
	handlerPemeriksaanEkg "rme/internal/handler/pemeriksaan_ekg"
	handlerPemeriksaanFungsiOrgan "rme/internal/handler/pemeriksaan_fungsi_organ"
	handlerPemeriksaanLaboratorium "rme/internal/handler/pemeriksaan_laboratorium"
	handlerPemeriksaanPenunjangBedah "rme/internal/handler/pemeriksaan_penunjang_bedah"
	handlerPemeriksaanVital "rme/internal/handler/pemeriksaan_vital"
	handlerPoli "rme/internal/handler/poli"
	handlerRencanaRawat "rme/internal/handler/rencana_rawat"
	handlerRujukUlang "rme/internal/handler/rujuk_ulang"
	handlerTatalaksana "rme/internal/handler/tatalaksana"

	mysqlAnamnesis "rme/internal/mysql/anamnesis"
	mysqlAntrian "rme/internal/mysql/antrian"
	mysqlDiagnosis "rme/internal/mysql/diagnosis"
	mysqlDoctor "rme/internal/mysql/doctor"
	mysqlICD "rme/internal/mysql/icd"
	mysqlLokalisBedah "rme/internal/mysql/lokalis_bedah"
	mysqlPatient "rme/internal/mysql/patient"
	mysqlPemeriksaanEkg "rme/internal/mysql/pemeriksaan_ekg"
	mysqlPemeriksaanFungsiOrgan "rme/internal/mysql/pemeriksaan_fungsi_organ"
	mysqlPemeriksaanLaboratorium "rme/internal/mysql/pemeriksaan_laboratorium"
	mysqlPemeriksaanPenunjangBedah "rme/internal/mysql/pemeriksaan_penunjang_bedah"
	mysqlPemeriksaanVital "rme/internal/mysql/pemeriksaan_vital"
	mysqlPoli "rme/internal/mysql/poli"
	mysqlRencanaRawat "rme/internal/mysql/rencana_rawat"
	mysqlRujukUlang "rme/internal/mysql/rujuk_ulang"
	mysqlTatalaksana "rme/internal/mysql/tatalaksana"
	mysqlUser "rme/internal/mysql/user"

	repoAntrianInterface "rme/internal/repository/antrian"
	repoPoliInterface "rme/internal/repository/poli"
	router "rme/internal/router"
	usecaseAnamnesis "rme/internal/usecase/anamnesis"
	usecaseAntrian "rme/internal/usecase/antrian"
	usecaseAuth "rme/internal/usecase/auth"
	usecaseDiagnosis "rme/internal/usecase/diagnosis"
	usecaseDoctor "rme/internal/usecase/doctor"
	usecaseICD "rme/internal/usecase/icd"
	usecaseLokalisBedah "rme/internal/usecase/lokalis_bedah"
	usecasePatient "rme/internal/usecase/patient"
	usecasePemeriksaanEkg "rme/internal/usecase/pemeriksaan_ekg"
	usecasePemeriksaanFungsiOrgan "rme/internal/usecase/pemeriksaan_fungsi_organ"
	usecasePemeriksaanLaboratorium "rme/internal/usecase/pemeriksaan_laboratorium"
	usecasePemeriksaanPenunjangBedah "rme/internal/usecase/pemeriksaan_penunjang_bedah"
	usecasePemeriksaanVital "rme/internal/usecase/pemeriksaan_vital"
	usecasePoli "rme/internal/usecase/poli"
	usecaseRencanaRawat "rme/internal/usecase/rencana_rawat"
	usecaseRujukUlang "rme/internal/usecase/rujuk_ulang"
	usecaseTatalaksana "rme/internal/usecase/tatalaksana"
)

func main() {
	var httpPortFlag string
	flag.StringVar(&httpPortFlag, "http-port", "", "HTTP listen port (overrides APP_PORT/PORT env var)")
	flag.Parse()

	// Ambil konfigurasi DB dari environment (sesuai launch.json)
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || host == "" || port == "" || name == "" {
		log.Fatal("database environment variables (DB_USER, DB_HOST, DB_PORT, DB_NAME) must not be empty")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	// Inisialisasi GORM MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Wiring clean architecture untuk Anamnesis
	repoAnamnesis := mysqlAnamnesis.NewRepository(db)
	ucAnamnesis := usecaseAnamnesis.NewUsecase(repoAnamnesis)
	handlerAnam := handlerAnamnesis.NewHandler(ucAnamnesis)

	// Wiring clean architecture untuk Auth/Login
	repoUser := mysqlUser.NewRepository(db)
	ucAuth := usecaseAuth.NewUsecase(repoUser)
	handlerAuthHTTP := handlerAuth.NewHandler(ucAuth)

	// Wiring clean architecture untuk Patient
	repoPatient := mysqlPatient.NewRepository(db)
	ucPatient := usecasePatient.NewUsecase(repoPatient)
	handlerPatientHTTP := handlerPatient.NewHandler(ucPatient)

	// Wiring clean architecture untuk Doctor
	repoDoctor := mysqlDoctor.NewRepository(db)
	ucDoctor := usecaseDoctor.NewUsecase(repoDoctor)
	handlerDoctorHTTP := handlerDoctor.NewHandler(ucDoctor)

	// Wiring clean architecture untuk Poli
	repoPoli := mysqlPoli.NewRepositoryMySQL(db)
	var repoPoliIface repoPoliInterface.Repository = repoPoli
	ucPoli := usecasePoli.NewUsecase(repoPoliIface)
	handlerPoliHTTP := handlerPoli.NewHandler(ucPoli)

	// Wiring clean architecture untuk Antrian
	repoAntrian := mysqlAntrian.NewRepositoryMySQL(db)
	var repoAntrianIface repoAntrianInterface.Repository = repoAntrian
	ucAntrian := usecaseAntrian.NewUsecase(repoAntrianIface)
	handlerAntrianHTTP := handlerAntrian.NewHandler(ucAntrian)

	// Wiring pemeriksaan_vital
	repoPemeriksaanVital := mysqlPemeriksaanVital.NewRepository(db)
	ucPemeriksaanVital := usecasePemeriksaanVital.NewUsecase(repoPemeriksaanVital)
	handlerPemeriksaanVitalHTTP := handlerPemeriksaanVital.NewHandler(ucPemeriksaanVital)

	// Wiring pemeriksaan_laboratorium
	repoPemeriksaanLaboratorium := mysqlPemeriksaanLaboratorium.NewRepository(db)
	ucPemeriksaanLaboratorium := usecasePemeriksaanLaboratorium.NewUsecase(repoPemeriksaanLaboratorium)
	handlerPemeriksaanLaboratoriumHTTP := handlerPemeriksaanLaboratorium.NewHandler(ucPemeriksaanLaboratorium)

	// Wiring pemeriksaan_ekg
	repoPemeriksaanEkg := mysqlPemeriksaanEkg.NewRepository(db)
	ucPemeriksaanEkg := usecasePemeriksaanEkg.NewUsecase(repoPemeriksaanEkg)
	handlerPemeriksaanEkgHTTP := handlerPemeriksaanEkg.NewHandler(ucPemeriksaanEkg)

	// Wiring ICD
	repoICD := mysqlICD.NewRepository(db)
	ucICD := usecaseICD.NewUsecase(repoICD)
	handlerICDHTTP := handlerICD.NewHandler(ucICD)

	// Wiring tatalaksana
	repoTatalaksana := mysqlTatalaksana.NewRepository(db)
	ucTatalaksana := usecaseTatalaksana.NewUsecase(repoTatalaksana)
	handlerTatalaksanaHTTP := handlerTatalaksana.NewHandler(ucTatalaksana)

	// Wiring status lokalis bedah
	repoLokalisBedah := mysqlLokalisBedah.NewRepository(db)
	ucLokalisBedah := usecaseLokalisBedah.NewUsecase(repoLokalisBedah)
	handlerLokalisBedahHTTP := handlerLokalisBedah.NewHandler(ucLokalisBedah)

	// Wiring pemeriksaan fungsi organ
	repoPemeriksaanFungsiOrgan := mysqlPemeriksaanFungsiOrgan.NewRepository(db)
	ucPemeriksaanFungsiOrgan := usecasePemeriksaanFungsiOrgan.NewUsecase(repoPemeriksaanFungsiOrgan)
	handlerPemeriksaanFungsiOrganHTTP := handlerPemeriksaanFungsiOrgan.NewHandler(ucPemeriksaanFungsiOrgan)

	// Wiring pemeriksaan penunjang bedah (new)
	repoPemeriksaanPenunjangBedah := mysqlPemeriksaanPenunjangBedah.NewRepository(db)
	ucPemeriksaanPenunjangBedah := usecasePemeriksaanPenunjangBedah.NewUsecase(repoPemeriksaanPenunjangBedah)
	handlerPemeriksaanPenunjangBedahHTTP := handlerPemeriksaanPenunjangBedah.NewHandler(ucPemeriksaanPenunjangBedah)

	// Wiring diagnosis
	repoDiagnosis := mysqlDiagnosis.NewMySQLRepo(db)
	ucDiagnosis := usecaseDiagnosis.NewUsecase(repoDiagnosis)
	handlerDiagnosisHTTP := handlerDiagnosis.NewHandler(ucDiagnosis)
	repoRencanaRawat := mysqlRencanaRawat.NewRepository(db)
	ucRencanaRawat := usecaseRencanaRawat.NewUsecase(repoRencanaRawat)
	handlerRencanaRawatHTTP := handlerRencanaRawat.NewHandler(ucRencanaRawat)
	repoRujukUlang := mysqlRujukUlang.NewRepository(db)
	ucRujukUlang := usecaseRujukUlang.NewUsecase(repoRujukUlang)
	handlerRujukUlangHTTP := handlerRujukUlang.NewHandler(ucRujukUlang)
	r := gin.Default()
	r.Use(cors.Default())
	router.RegisterAnamnesisRoutes(r, handlerAnam)
	router.RegisterAuthRoutes(r, handlerAuthHTTP)
	router.RegisterPatientRoutes(r, handlerPatientHTTP)
	router.RegisterDoctorRoutes(r, handlerDoctorHTTP)
	router.RegisterPoliRoutes(r, handlerPoliHTTP)
	router.RegisterAntrianRoutes(r, handlerAntrianHTTP)
	router.RegisterPemeriksaanVitalRoutes(r, handlerPemeriksaanVitalHTTP)
	router.RegisterPemeriksaanLaboratoriumRoutes(r, handlerPemeriksaanLaboratoriumHTTP)
	router.RegisterPemeriksaanEkgRoutes(r, handlerPemeriksaanEkgHTTP)
	router.RegisterICDRoutes(r, handlerICDHTTP)
	router.RegisterTatalaksanaRoutes(r, handlerTatalaksanaHTTP)
	router.RegisterRujukUlangRoutes(r, handlerRujukUlangHTTP)
	router.RegisterRencanaRawatRoutes(r, handlerRencanaRawatHTTP)
	router.RegisterDiagnosisRoutes(r, handlerDiagnosisHTTP)
	router.RegisterLokalisBedahRoutes(r, handlerLokalisBedahHTTP)
	router.RegisterPemeriksaanFungsiOrganRoutes(r, handlerPemeriksaanFungsiOrganHTTP)
	router.RegisterPemeriksaanPenunjangBedahRoutes(r, handlerPemeriksaanPenunjangBedahHTTP)

	httpPort := strings.TrimSpace(httpPortFlag)
	if httpPort == "" {
		httpPort = strings.TrimSpace(os.Getenv("APP_PORT"))
	}
	if httpPort == "" {
		httpPort = strings.TrimSpace(os.Getenv("PORT"))
	}
	if httpPort == "" {
		httpPort = "8080"
	}

	listenAddr := httpPort
	if !strings.Contains(listenAddr, ":") {
		listenAddr = ":" + listenAddr
	}
	log.Printf("listening and serving HTTP on %s", listenAddr)
	if err := r.Run(listenAddr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
