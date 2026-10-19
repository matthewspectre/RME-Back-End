package rencanarawat

import entity "rme/internal/entity/rencana_rawat"

type Repository interface {
	Create(data *entity.RencanaRawat) error
	GetByID(id int) (*entity.RencanaRawat, error)
	GetAll(idPasien *int, idDokter *int) ([]*entity.RencanaRawat, error)
	Update(id int, updates map[string]interface{}) error
}
