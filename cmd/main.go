package main

import (
	"github.com/SabienNguyen/military-logistics-planner/internal/db"
	"github.com/SabienNguyen/military-logistics-planner/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 🔌 Initialize the DB connection
	database := db.Init()

	// 🌐 Create Gin router
	r := gin.Default()

	// 🧩 Register your routes, passing in the DB connection
	handlers.RegisterZoneRoutes(r, database)
	handlers.RegisterAssignmentRoutes(r, database)
	handlers.RegisterResourceRoutes(r, database)
	handlers.RegisterLogRoutes(r, database)
	handlers.RegisterAuthRoutes(r, database)

	// 🚀 Run the server
	r.Run(":8080")
}
