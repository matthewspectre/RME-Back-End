package rencanarawat

import (
	"net/http"
	"strconv"
	"time"

	entity "rme/internal/entity/rencana_rawat"
	usecase "rme/internal/usecase/rencana_rawat"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var req RencanaRawatCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.IDPasien <= 0 || req.IDDokter <= 0 || req.RencanaRawat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idPasien, idDokter, and rencana_rawat are required"})
		return
	}

	const layout = "2006-01-02 15:04:05"
	dateMake := time.Now()
	if req.DateMake != "" {
		parsed, err := time.Parse(layout, req.DateMake)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dateMake format"})
			return
		}
		dateMake = parsed
	}
	dateUpdate := dateMake
	if req.DateUpdate != "" {
		parsed, err := time.Parse(layout, req.DateUpdate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dateUpdate format"})
			return
		}
		dateUpdate = parsed
	}

	data := &entity.RencanaRawat{
		IDPasien:     req.IDPasien,
		IDDokter:     req.IDDokter,
		RencanaRawat: req.RencanaRawat,
		DateMake:     dateMake,
		DateUpdate:   dateUpdate,
		Visible:      1,
	}
	if err := h.uc.Create(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "rencana rawat created"})
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	data, err := h.uc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, toResponse(data))
}

func (h *Handler) GetAll(c *gin.Context) {
	idPasien, err := optionalID(c.Query("idPasien"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idPasien"})
		return
	}
	idDokter, err := optionalID(c.Query("idDokter"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idDokter"})
		return
	}
	list, err := h.uc.GetAll(idPasien, idDokter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := make([]RencanaRawatResponse, 0, len(list))
	for _, item := range list {
		response = append(response, toResponse(item))
	}
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if value, ok := payload["rencana_rawat"]; ok {
		text, valid := value.(string)
		if !valid || text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "rencana_rawat must be a non-empty string"})
			return
		}
		updates["rencana_rawat"] = text
	}
	if value, ok := payload["id_dokter"]; ok {
		doctorID, valid := numericID(value)
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id_dokter"})
			return
		}
		updates["id_dokter"] = doctorID
	}
	if value, ok := payload["date_update"]; ok {
		dateString, valid := value.(string)
		if !valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_update format"})
			return
		}
		const layout = "2006-01-02 15:04:05"
		dateUpdate, err := time.Parse(layout, dateString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_update format"})
			return
		}
		updates["date_update"] = dateUpdate
	}
	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no updatable fields provided"})
		return
	}
	if _, hasDateUpdate := updates["date_update"]; !hasDateUpdate {
		updates["date_update"] = time.Now()
	}
	if err := h.uc.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *Handler) Hide(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.uc.Update(id, map[string]interface{}{"visible": 0, "date_update": time.Now()}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "hidden"})
}

func optionalID(value string) (*int, error) {
	if value == "" {
		return nil, nil
	}
	id, err := strconv.Atoi(value)
	if err != nil || id <= 0 {
		return nil, strconv.ErrSyntax
	}
	return &id, nil
}

func numericID(value interface{}) (int, bool) {
	switch id := value.(type) {
	case float64:
		if id > 0 && id == float64(int(id)) {
			return int(id), true
		}
	case string:
		parsed, err := strconv.Atoi(id)
		if err == nil && parsed > 0 {
			return parsed, true
		}
	}
	return 0, false
}

func toResponse(data *entity.RencanaRawat) RencanaRawatResponse {
	return RencanaRawatResponse{
		ID:           data.ID,
		IDPasien:     data.IDPasien,
		IDDokter:     data.IDDokter,
		RencanaRawat: data.RencanaRawat,
		DateMake:     data.DateMake,
		DateUpdate:   data.DateUpdate,
		Visible:      data.Visible,
		NamaPasien:   data.NamaPasien,
		NamaDokter:   data.NamaDokter,
	}
}
