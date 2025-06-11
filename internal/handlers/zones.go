package handlers

import (
	"net/http"

	"github.com/SabienNguyen/military-logistics-planner/internal/auth"
	"github.com/SabienNguyen/military-logistics-planner/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterZoneRoutes sets up the /zones endpoints
func RegisterZoneRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/zones", auth.RequireRole("admin"), func(c *gin.Context) {
		var zone models.Zone
		if err := c.ShouldBindJSON(&zone); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&zone).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create zone"})
			return
		}

		c.JSON(http.StatusCreated, zone)
	})

	r.GET("/zones", func(c *gin.Context) {
		var zones []models.Zone
		if err := db.Find(&zones).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch zones"})
			return
		}
		c.JSON(http.StatusOK, zones)
	})

	r.GET("/zones/:id", func(c *gin.Context) {
		var zone models.Zone
		id := c.Param("id")

		if err := db.First(&zone, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "zone not found"})
			return
		}

		c.JSON(http.StatusOK, zone)
	})
}
