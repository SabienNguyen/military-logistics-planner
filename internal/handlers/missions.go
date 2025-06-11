package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/SabienNguyen/military-logistics-planner/internal/auth"
	"github.com/SabienNguyen/military-logistics-planner/internal/models"
)

func RegisterMissionRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/missions")
	api.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Public (read-only, anyone logged in)
	api.GET("", auth.RequireRole("admin", "officer", "viewer"), GetMissions)
	api.GET("/:id", auth.RequireRole("admin", "officer", "viewer"), GetMissionByID)

	// Create, Update, Delete – Officers and Admins only
	api.POST("", auth.RequireRole("admin", "officer"), CreateMission)
	api.PATCH("/:id", auth.RequireRole("admin", "officer"), UpdateMission)
	api.DELETE("/:id", auth.RequireRole("admin"), DeleteMission)

	// Assign/remove resources – Officers and Admins only
	// api.POST("/:id/assign", auth.RequireRole("admin", "officer"), AssignResourcesToMission)
	// api.POST("/:id/remove", auth.RequireRole("admin", "officer"), RemoveResourcesFromMission)
}

// POST /api/missions
func CreateMission(c *gin.Context) {
	var input struct {
		Name        string    `json:"name" binding:"required"`
		Description string    `json:"description"`
		StartDate   time.Time `json:"start_date" binding:"required"`
		EndDate     time.Time `json:"end_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mission := models.Mission{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Description: input.Description,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
		Status:      "planned",
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&mission).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mission)
}

// GET /api/missions
func GetMissions(c *gin.Context) {
	var missions []models.Mission
	db := c.MustGet("db").(*gorm.DB)

	status := c.Query("status")
	query := db.Preload("AssignedResource").Order("created_at desc")
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Find(&missions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get missions"})
		return
	}

	c.JSON(http.StatusOK, missions)
}

// GET /api/missions/:id
func GetMissionByID(c *gin.Context) {
	id := c.Param("id")
	var mission models.Mission

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Preload("AssignedResource").First(&mission, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
		return
	}

	c.JSON(http.StatusOK, mission)
}

// PATCH /api/missions/:id
func UpdateMission(c *gin.Context) {
	id := c.Param("id")
	var mission models.Mission

	db := c.MustGet("db").(*gorm.DB)
	if err := db.First(&mission, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
		return
	}

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&mission).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update mission"})
		return
	}

	db.Preload("AssignedResource").First(&mission, "id = ?", id)
	c.JSON(http.StatusOK, mission)
}

// DELETE /api/missions/:id
func DeleteMission(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Delete(&models.Mission{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete mission"})
		return
	}

	c.Status(http.StatusNoContent)
}

// POST /api/missions/:id/assign
// func AssignResourcesToMission(c *gin.Context) {
// 	id := c.Param("id")
// 	var input struct {
// 		ResourceIDs []uuid.UUID `json:"resource_ids"`
// 	}
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db := c.MustGet("db").(*gorm.DB)
// 	var mission models.Mission
// 	if err := db.Preload("AssignedResource").First(&mission, "id = ?", id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
// 		return
// 	}

// 	var resources []models.Resource
// 	if err := db.Where("id IN ?", input.ResourceIDs).Find(&resources).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch resources"})
// 		return
// 	}

// 	if err := db.Model(&mission).Association("AssignedResource").Append(&resources); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign resources"})
// 		return
// 	}

// 	db.Preload("AssignedResource").First(&mission)
// 	c.JSON(http.StatusOK, mission.AssignedResource)
// }

// POST /api/missions/:id/remove
// func RemoveResourcesFromMission(c *gin.Context) {
// 	id := c.Param("id")
// 	var input struct {
// 		ResourceIDs []uuid.UUID `json:"resource_ids"`
// 	}
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db := c.MustGet("db").(*gorm.DB)
// 	var mission models.Mission
// 	if err := db.Preload("AssignedResource").First(&mission, "id = ?", id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
// 		return
// 	}

// 	var resources []models.Resource
// 	if err := db.Where("id IN ?", input.ResourceIDs).Find(&resources).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch resources"})
// 		return
// 	}

// 	if err := db.Model(&mission).Association("AssignedResource").Delete(&resources); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove resources"})
// 		return
// 	}

// 	db.Preload("AssignedResource").First(&mission)
// 	c.JSON(http.StatusOK, mission.AssignedResource)
// }
