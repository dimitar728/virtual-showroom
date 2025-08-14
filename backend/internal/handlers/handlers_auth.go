package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RegisterPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"` // optional: admin or user (you might restrict this in production)
}

func RegisterHandler(db *gorm.DB, cfg Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p RegisterPayload
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// prevent creating admin via public endpoint in production
		role := "user"
		if p.Role == "admin" {
			role = "user"
		}
		hash, err := HashPassword(p.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash"})
			return
		}
		u := User{
			ID:           uuid.New(),
			Email:        p.Email,
			PasswordHash: hash,
			Role:         role,
		}
		if err := db.Create(&u).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists or invalid"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": u.ID, "email": u.Email})
	}
}

type LoginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(db *gorm.DB, cfg Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p LoginPayload
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var user User
		if err := db.First(&user, "email = ?", p.Email).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		if user.Suspended {
			c.JSON(http.StatusForbidden, gin.H{"error": "account suspended"})
			return
		}
		if !CheckPasswordHash(p.Password, user.PasswordHash) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		token, err := GenerateJWT(user.ID, user.Role, cfg.JWTSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func MeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		v, _ := c.Get("currentUser")
		user := v.(User)
		// hide password
		c.JSON(http.StatusOK, gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"role":       user.Role,
			"suspended":  user.Suspended,
			"created_at": user.CreatedAt,
		})
	}
}
