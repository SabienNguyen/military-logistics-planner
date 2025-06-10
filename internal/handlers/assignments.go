package handlers

import (
	"net/http"

	"github.com/SabienNguyen/military-logistics-planner/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AssignRequest struct {
	ResourceID uint   `json:"resource_id"`
	ToZoneID   uint   `json:"to_zone_id"`
	Note       string `json:"note"`
}

func RegisterAssignmentRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/assign", func(c *gin.Context) {
		var req AssignRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var resource models.Resource
		if err := db.First(&resource, req.ResourceID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "resource not found"})
			return
		}

		fromZoneID := resource.ZoneID // track where it's coming from
		resource.ZoneID = req.ToZoneID

		if err := db.Save(&resource).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to reassign resource"})
			return
		}

		log := models.MovementLog{
			ResourceID: resource.ID,
			FromZoneID: &fromZoneID,
			ToZoneID:   req.ToZoneID,
			Note:       req.Note,
		}

		if err := db.Create(&log).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to log movement"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "resource reassigned", "movement_log": log})
	})
}
