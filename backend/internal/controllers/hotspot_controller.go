package controllers

import (
	"net/http"

	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/dimitar728/virtual-showroom/backend/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateHotspotDTO struct {
	Label       string  `json:"label" binding:"required"`
	Description string  `json:"description"`
	PosX        float64 `json:"pos_x" binding:"required"`
	PosY        float64 `json:"pos_y" binding:"required"`
	PosZ        float64 `json:"pos_z" binding:"required"`
	MediaURL    string  `json:"media_url"`
	LinkURL     string  `json:"link_url"`
}

type UpdateHotspotDTO struct {
	Label       *string  `json:"label"`
	Description *string  `json:"description"`
	PosX        *float64 `json:"pos_x"`
	PosY        *float64 `json:"pos_y"`
	PosZ        *float64 `json:"pos_z"`
	MediaURL    *string  `json:"media_url"`
	LinkURL     *string  `json:"link_url"`
}

// GET /api/showrooms/:id/hotspots
func ListHotspots(c *gin.Context) {
	showroomID := c.Param("id")
	var hotspots []models.Hotspot
	if err := database.DB.Where("showroom_id = ?", showroomID).Find(&hotspots).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch hotspots"})
		return
	}
	c.JSON(http.StatusOK, hotspots)
}

// POST /api/showrooms/:id/hotspots (admin)
func CreateHotspot(c *gin.Context) {
	showroomID := c.Param("id")
	var body CreateHotspotDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := uuid.Parse(showroomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid showroom id"})
		return
	}
	h := models.Hotspot{
		ShowroomID:  id,
		Label:       body.Label,
		Description: body.Description,
		PosX:        body.PosX,
		PosY:        body.PosY,
		PosZ:        body.PosZ,
		MediaURL:    body.MediaURL,
		LinkURL:     body.LinkURL,
	}
	if err := database.DB.Create(&h).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create hotspot"})
		return
	}
	c.JSON(http.StatusCreated, h)
}

// PATCH /api/hotspots/:hid (admin)
func UpdateHotspot(c *gin.Context) {
	hid := c.Param("hid")
	var body UpdateHotspotDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var h models.Hotspot
	if err := database.DB.First(&h, "id = ?", hid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "hotspot not found"})
		return
	}

	if body.Label != nil {
		h.Label = *body.Label
	}
	if body.Description != nil {
		h.Description = *body.Description
	}
	if body.PosX != nil {
		h.PosX = *body.PosX
	}
	if body.PosY != nil {
		h.PosY = *body.PosY
	}
	if body.PosZ != nil {
		h.PosZ = *body.PosZ
	}
	if body.MediaURL != nil {
		h.MediaURL = *body.MediaURL
	}
	if body.LinkURL != nil {
		h.LinkURL = *body.LinkURL
	}

	if err := database.DB.Save(&h).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update hotspot"})
		return
	}
	c.JSON(http.StatusOK, h)
}

// DELETE /api/hotspots/:hid (admin)
func DeleteHotspot(c *gin.Context) {
	hid := c.Param("hid")
	if err := database.DB.Delete(&models.Hotspot{}, "id = ?", hid).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete hotspot"})
		return
	}
	c.Status(http.StatusNoContent)
}
