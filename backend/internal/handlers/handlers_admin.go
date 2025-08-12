package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET /api/admin/users
func AdminListUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []User
		db.Find(&users)
		// hide passwordhash
		out := make([]map[string]interface{}, 0, len(users))
		for _, u := range users {
			out = append(out, map[string]interface{}{
				"id":         u.ID,
				"email":      u.Email,
				"role":       u.Role,
				"suspended":  u.Suspended,
				"created_at": u.CreatedAt,
			})
		}
		c.JSON(http.StatusOK, out)
	}
}

type UserPatchPayload struct {
	Action string `json:"action" binding:"required"` // suspend | reactivate
}

func AdminPatchUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var payload UserPatchPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user User
		if err := db.First(&user, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		switch payload.Action {
		case "suspend":
			user.Suspended = true
			db.Save(&user)
			c.JSON(http.StatusOK, gin.H{"status": "suspended"})
			return
		case "reactivate":
			user.Suspended = false
			db.Save(&user)
			c.JSON(http.StatusOK, gin.H{"status": "reactivated"})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid action"})
			return
		}
	}
}

func AdminDeleteUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// Prevent deleting self maybe - optional
		if err := db.Delete(&User{}, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
