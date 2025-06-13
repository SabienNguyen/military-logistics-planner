package handlers

import (
	"net/http"

	"github.com/SabienNguyen/military-logistics-planner/internal/auth"
	"github.com/SabienNguyen/military-logistics-planner/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterResourceRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/resources", auth.RequireRole("admin", "officer"), func(c *gin.Context) {
		var resource models.Resource
		if err := c.ShouldBindJSON(&resource); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&resource).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create resource"})
			return
		}
		c.JSON(http.StatusCreated, resource)
	})
	r.GET("/zones/:id/resources", auth.RequireRole("viewer", "officer", "admin"), func(c *gin.Context) {
		zoneID := c.Param("id")

		var resources []models.Resource
		if err := db.Where("zone_id = ?", zoneID).Find(&resources).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch resources"})
			return
		}

		c.JSON(http.StatusOK, resources)
	})

}
