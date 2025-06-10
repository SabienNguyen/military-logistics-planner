package handlers

import (
	"net/http"

	"github.com/SabienNguyen/military-logistics-planner/internal/auth"
	"github.com/SabienNguyen/military-logistics-planner/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		var user models.User
		if err := db.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		token, err := auth.GenerateToken(user.ID, user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})
}
