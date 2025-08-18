package controllers

import (

	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/dimitar728/virtual-showroom/backend/internal/database"
	"github.com/dimitar728/virtual-showroom/backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


// --- Helpers ---
func validateModelFile(file *multipart.FileHeader) error {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".glb" && ext != ".gltf" {
		return fmt.Errorf("invalid file type: only .glb/.gltf allowed")
	}
	if file.Size > 50*1024*1024 { // 50MB limit
		return fmt.Errorf("file too large (max 50MB)")
	}
	return nil
}

// GET /api/showrooms
func CreateShowroom(c *fiber.Ctx) error {

// GET /api/showrooms
func CreateShowroom(c *fiber.Ctx) error {
	if err := utils.ValidateModelPath(filename); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}


	// Parse text fields
	name := c.FormValue("name")
	description := c.FormValue("description")
	capacity := c.FormValue("capacity")

	if name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name is required"})
	}

	// Handle file upload
	file, err := c.FormFile("model")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "model file required"})
	}
	if err := validateModelFile(file); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Save file locally (later → S3 or Cloudinary)
	filename := fmt.Sprintf("uploads/models/%d_%s", time.Now().Unix(), file.Filename)
	if err := c.SaveFile(file, filename); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to save model"})
	}

	// Build showroom record
	adminID := c.Locals("userID").(uuid.UUID)
	showroom := models.Showroom{
		Name:        name,
		Description: description,
		ModelPath:   filename,
		CreatedBy:   adminID,
		CreatedAt:   time.Now(),
	}

	if capacity != "" {
		fmt.Sscan(capacity, &showroom.Capacity)
	}

	if err := database.DB.Create(&showroom).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create showroom"})
	}
	return c.Status(201).JSON(showroom)


	if err := utils.ValidateModelPath(filename); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

}

func GetShowrooms(c *fiber.Ctx) error {
	var showrooms []models.Showroom
	if err := database.DB.Find(&showrooms).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch showrooms"})
	}
	return c.JSON(showrooms)

}

// GET /api/showrooms/:id
func GetShowroomByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var showroom models.Showroom

	if err := database.DB.First(&showroom, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "showroom not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch showroom"})
	}
	return c.JSON(showroom)
}


// PATCH /api/showrooms/:id (admin only)
func UpdateShowroom(c *fiber.Ctx) error {

// POST /api/showrooms (admin only)
func CreateShowroom(c *fiber.Ctx) error {
	var body models.Showroom
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if body.Name == "" || body.ModelPath == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name and model_path required"})
	}

	// Extract admin ID from context (set by middleware)
	adminID := c.Locals("userID").(uuid.UUID)
	body.CreatedBy = adminID

	if err := database.DB.Create(&body).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create showroom"})
	}
	return c.Status(201).JSON(body)

}

// PATCH /api/showrooms/:id (admin only)
func UpdateShowroom(c *fiber.Ctx) error {
	if err := utils.ValidateModelPath(filename); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
}


	id := c.Params("id")
	var showroom models.Showroom

	if err := database.DB.First(&showroom, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"error": "showroom not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "db error"})
	}

	// Parse updates
	name := c.FormValue("name")
	description := c.FormValue("description")
	capacity := c.FormValue("capacity")

	if name != "" {
		showroom.Name = name
	}
	if description != "" {
		showroom.Description = description
	}
	if capacity != "" {
		fmt.Sscan(capacity, &showroom.Capacity)
	}

	// Optional file update
	file, _ := c.FormFile("model")
	if file != nil {
		if err := validateModelFile(file); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		filename := fmt.Sprintf("uploads/models/%d_%s", time.Now().Unix(), file.Filename)
		if err := c.SaveFile(file, filename); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to save model"})
		}
		showroom.ModelPath = filename
	}

	if err := database.DB.Save(&showroom).Error; err != nil {

		return c.Status(500).JSON(fiber.Map{"error": "failed to update showroom"})
	}
	return c.JSON(showroom)

	if err := utils.ValidateModelPath(filename); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
}



		return c.Status(404).JSON(fiber.Map{"error": "showroom not found"})
	}

	var body map[string]interface{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := database.DB.Model(&showroom).Updates(body).Error; err != nil {

		return c.Status(500).JSON(fiber.Map{"error": "failed to update showroom"})
	}
	return c.JSON(showroom)
}



// DELETE /api/showrooms/:id (admin only)
func DeleteShowroom(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Showroom{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete showroom"})
	}
	return c.SendStatus(204)
}


func UploadModel(c *gin.Context) {
	id := c.Param("id")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing file"})
		return
	}

	// save file
	savePath := "uploads/models/" + file.Filename
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	if err := database.DB.Model(&models.Showroom{}).
		Where("id = ?", id).
		Update("model_path", savePath).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update showroom"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"model_path": savePath})
}

