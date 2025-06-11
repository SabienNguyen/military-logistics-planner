package main

import (
	"os"

	"github.com/SabienNguyen/military-logistics-planner/internal/db"
	"github.com/SabienNguyen/military-logistics-planner/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 🔌 Initialize the DB connection
	database := db.Init()

	// 🌐 Create Gin router
	r := gin.Default()

	//set jwt secret
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// 🧩 Register your routes, passing in the DB connection
	handlers.RegisterZoneRoutes(r, database)
	handlers.RegisterAssignmentRoutes(r, database)
	handlers.RegisterResourceRoutes(r, database)
	handlers.RegisterLogRoutes(r, database)
	handlers.RegisterAuthRoutes(r, database, jwtSecret)
	handlers.RegisterMissionRoutes(r, database)

	// 🚀 Run the server
	r.Run(":8080")
}
