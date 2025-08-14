package handlers

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// List public showrooms
func ListShowroomsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var s []Showroom
		db.Find(&s)
		c.JSON(http.StatusOK, s)
	}
}

func GetShowroomHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var s Showroom
		if err := db.First(&s, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusOK, s)
	}
}

// UploadModelMiddleware accepts optional file in "model" form key and validates it
func UploadModelMiddleware(cfg Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if multipart/form-data with file
		form, err := c.MultipartForm()
		if err != nil {
			// no file; continue (update might not include file)
			c.Next()
			return
		}
		files := form.File["model"]
		if len(files) == 0 {
			c.Next()
			return
		}
		// Only accept first file
		f := files[0]
		if !IsValidModelFileName(f.Filename) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid model file (only .glb/.gltf allowed)"})
			return
		}
		// size check
		maxBytes := int64(cfg.MaxUploadMB) * 1024 * 1024
		if f.Size > maxBytes {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file too large"})
			return
		}
		// Save to uploads dir with uuid name
		dstName := uuid.New().String() + filepath.Ext(f.Filename)
		dstPath := filepath.Join(cfg.UploadDir, dstName)
		if err := c.SaveUploadedFile(f, dstPath); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to store file"})
			return
		}
		// Make path available to handlers
		c.Set("uploadedModelPath", dstPath)
		c.Next()
	}
}

type CreateShowroomPayload struct {
	Name        string `form:"name" binding:"required"`
	Description string `form:"description"`
	Capacity    int    `form:"capacity"`
}

func CreateShowroomHandler(db *gorm.DB, cfg Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p CreateShowroomPayload
		if err := c.ShouldBind(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// get current user
		v, _ := c.Get("currentUser")
		user := v.(User)

		modelPathIfAny, _ := c.Get("uploadedModelPath")
		modelPath := ""
		if mp, ok := modelPathIfAny.(string); ok {
			modelPath = mp
		}

		s := Showroom{
			ID:          uuid.New(),
			Name:        p.Name,
			Description: p.Description,
			ModelPath:   modelPath,
			Capacity:    p.Capacity,
			CreatedBy:   user.ID,
			CreatedAt:   time.Now(),
		}
		if err := db.Create(&s).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create showroom"})
			return
		}
		c.JSON(http.StatusCreated, s)
	}
}

type UpdateShowroomPayload struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Capacity    *int   `form:"capacity"`
}

func UpdateShowroomHandler(db *gorm.DB, cfg Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var s Showroom
		if err := db.First(&s, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		var p UpdateShowroomPayload
		if err := c.ShouldBind(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if p.Name != "" {
			s.Name = p.Name
		}
		if p.Description != "" {
			s.Description = p.Description
		}
		if p.Capacity != nil {
			s.Capacity = *p.Capacity
		}
		// handle uploaded file
		if mpv, ok := c.Get("uploadedModelPath"); ok {
			if pathStr, ok2 := mpv.(string); ok2 {
				s.ModelPath = pathStr
			}
		}
		if err := db.Save(&s).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update"})
			return
		}
		c.JSON(http.StatusOK, s)
	}
}

func DeleteShowroomHandler(db *gorm.DB, cfg Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&Showroom{}, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
