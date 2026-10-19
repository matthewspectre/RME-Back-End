package rencanarawat

import "time"

type RencanaRawatModel struct {
	ID           int       `gorm:"primaryKey;autoIncrement;column:id_rencana_rawat"`
	IDPasien     int       `gorm:"column:id_pasien"`
	IDDokter     int       `gorm:"column:id_dokter"`
	RencanaRawat string    `gorm:"column:rencana_rawat"`
	DateMake     time.Time `gorm:"column:date_make"`
	DateUpdate   time.Time `gorm:"column:date_update"`
	Visible      int       `gorm:"column:visible"`
}

func (RencanaRawatModel) TableName() string {
	return "rencana_rawat"
}
