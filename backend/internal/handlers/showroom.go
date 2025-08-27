package handlers

import (
	"net/http"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShowroomHandler struct{ DB *gorm.DB }

func (h ShowroomHandler) List(c *gin.Context) {
	var s []models.Showroom
	h.DB.Order("created_at desc").Find(&s)
	c.JSON(http.StatusOK, s)
}

func (h ShowroomHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var s models.Showroom
	if err := h.DB.First(&s, "id = ?", uuid.MustParse(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, s)
}

func (h ShowroomHandler) Create(c *gin.Context) {
	var s models.Showroom
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.CreatedBy = uuid.MustParse(c.GetString("uid"))
	if err := h.DB.Create(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create"})
		return
	}
	c.JSON(http.StatusCreated, s)
}

func (h ShowroomHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var s models.Showroom
	if err := h.DB.First(&s, "id = ?", uuid.MustParse(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	var req map[string]any
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Model(&s).Updates(req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, s)
}

func (h ShowroomHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Showroom{}, "id = ?", uuid.MustParse(id)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.Status(http.StatusNoContent)
}
