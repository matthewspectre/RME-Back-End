package rencanarawat

import (
	"context"
	"time"

	entity "rme/internal/entity/rencana_rawat"
	model "rme/internal/model/rencana_rawat"
	repo "rme/internal/repository/rencana_rawat"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toModel(e *entity.RencanaRawat) *model.RencanaRawatModel {
	if e == nil {
		return nil
	}
	return &model.RencanaRawatModel{
		ID:           e.ID,
		IDPasien:     e.IDPasien,
		IDDokter:     e.IDDokter,
		RencanaRawat: e.RencanaRawat,
		DateMake:     e.DateMake,
		DateUpdate:   e.DateUpdate,
		Visible:      e.Visible,
	}
}

type joinedRow struct {
	ID           int       `gorm:"column:id_rencana_rawat"`
	IDPasien     int       `gorm:"column:id_pasien"`
	IDDokter     int       `gorm:"column:id_dokter"`
	RencanaRawat string    `gorm:"column:rencana_rawat"`
	DateMake     time.Time `gorm:"column:date_make"`
	DateUpdate   time.Time `gorm:"column:date_update"`
	Visible      int       `gorm:"column:visible"`
	NamaPasien   string    `gorm:"column:nama_pasien"`
	NamaDokter   string    `gorm:"column:nama_dokter"`
}

func rowToEntity(row *joinedRow) *entity.RencanaRawat {
	if row == nil {
		return nil
	}
	return &entity.RencanaRawat{
		ID:           row.ID,
		IDPasien:     row.IDPasien,
		IDDokter:     row.IDDokter,
		RencanaRawat: row.RencanaRawat,
		DateMake:     row.DateMake,
		DateUpdate:   row.DateUpdate,
		Visible:      row.Visible,
		NamaPasien:   row.NamaPasien,
		NamaDokter:   row.NamaDokter,
	}
}

func (r *RepositoryMySQL) selectWithNames() *gorm.DB {
	return r.db.Raw(`SELECT rr.*, p.name AS nama_pasien, d.nama_dokter AS nama_dokter
		FROM rencana_rawat rr
		LEFT JOIN patients p ON p.id = rr.id_pasien
		LEFT JOIN dokter d ON d.id = rr.id_dokter`)
}

func (r *RepositoryMySQL) Create(data *entity.RencanaRawat) error {
	return r.db.Create(toModel(data)).Error
}

func (r *RepositoryMySQL) GetByID(id int) (*entity.RencanaRawat, error) {
	var row joinedRow
	query := r.selectWithNames().Where("rr.id_rencana_rawat = ?", id)
	if err := query.WithContext(context.Background()).Scan(&row).Error; err != nil {
		return nil, err
	}
	if row.ID == 0 {
		return nil, nil
	}
	return rowToEntity(&row), nil
}

func (r *RepositoryMySQL) GetAll(idPasien *int, idDokter *int) ([]*entity.RencanaRawat, error) {
	query := r.selectWithNames().Where("rr.visible = ?", 1)
	if idPasien != nil {
		query = query.Where("rr.id_pasien = ?", *idPasien)
	}
	if idDokter != nil {
		query = query.Where("rr.id_dokter = ?", *idDokter)
	}
	query = query.Order("rr.date_make DESC")

	var rows []joinedRow
	if err := query.WithContext(context.Background()).Scan(&rows).Error; err != nil {
		return nil, err
	}
	result := make([]*entity.RencanaRawat, 0, len(rows))
	for i := range rows {
		result = append(result, rowToEntity(&rows[i]))
	}
	return result, nil
}

func (r *RepositoryMySQL) Update(id int, updates map[string]interface{}) error {
	if updates == nil || len(updates) == 0 {
		return nil
	}
	return r.db.WithContext(context.Background()).
		Model(&model.RencanaRawatModel{}).
		Where("id_rencana_rawat = ?", id).
		Updates(updates).Error
}
