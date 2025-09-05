package controllers

import (
	"net/http"

	"github.com/ajonesb/user-management/backend/internal/models"
	"github.com/ajonesb/user-management/backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ShowroomController struct {
	service *services.ShowroomService
}

func NewShowroomController(service *services.ShowroomService) *ShowroomController {
	return &ShowroomController{service: service}
}

func (c *ShowroomController) Create(ctx *gin.Context) {
	var input models.Showroom
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Extract user_id from JWT claims
	userIDVal, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}
	userIDStr, ok := userIDVal.(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
		return
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format"})
		return
	}
	if err := c.service.Create(&input, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, input)
}

func (c *ShowroomController) GetAll(ctx *gin.Context) {
	showrooms, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, showrooms)
}

func (c *ShowroomController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	showroom, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Showroom not found"})
		return
	}
	ctx.JSON(http.StatusOK, showroom)
}

func (c *ShowroomController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var showroom models.Showroom
	if err := ctx.ShouldBindJSON(&showroom); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid showroom ID"})
		return
	}
	showroom.ID = uid
	if err := c.service.Update(&showroom); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, showroom)
}

func (c *ShowroomController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Showroom deleted"})
}
