package handlers

import (
	"net/http"

	"github.com/SabienNguyen/military-logistics-planner/internal/auth"

	"github.com/SabienNguyen/military-logistics-planner/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterLogRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/logs", auth.RequireRole("officer", "admin", "viewer"), func(c *gin.Context) {
		var logs []models.MovementLog

		if err := db.Order("created_at desc").Find(&logs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve logs"})
			return
		}

		c.JSON(http.StatusOK, logs)
	})
}
