package rencanarawat

import "time"

type RencanaRawatCreateRequest struct {
	IDPasien     int    `json:"idPasien"`
	IDDokter     int    `json:"idDokter"`
	RencanaRawat string `json:"rencana_rawat"`
	DateMake     string `json:"dateMake"`
	DateUpdate   string `json:"dateUpdate"`
}

type RencanaRawatResponse struct {
	ID           int       `json:"id_rencana_rawat"`
	IDPasien     int       `json:"id_pasien"`
	IDDokter     int       `json:"id_dokter"`
	RencanaRawat string    `json:"rencana_rawat"`
	DateMake     time.Time `json:"date_make"`
	DateUpdate   time.Time `json:"date_update"`
	Visible      int       `json:"visible"`
	NamaPasien   string    `json:"nama_pasien"`
	NamaDokter   string    `json:"nama_dokter"`
}
