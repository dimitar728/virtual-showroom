package controllers

import (
	"net/http"

	"github.com/ajonesb/user-management/backend/internal/models"
	"github.com/ajonesb/user-management/backend/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) ListUsers(ctx *gin.Context) {
	users, err := c.userService.ListUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user.ID = uid
	if err := c.userService.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) SuspendUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.userService.SuspendUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to suspend user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User suspended"})
}

func (c *UserController) ReactivateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.userService.ReactivateUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reactivate user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User reactivated"})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.userService.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
