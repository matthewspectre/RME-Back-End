package rencanarawat

import "time"

type RencanaRawat struct {
	ID           int
	IDPasien     int
	IDDokter     int
	RencanaRawat string
	DateMake     time.Time
	DateUpdate   time.Time
	Visible      int
	NamaPasien   string
	NamaDokter   string
}
