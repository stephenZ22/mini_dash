package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/projects"
	"github.com/stephenZ22/mini_dash/internal/users"
	"github.com/stephenZ22/mini_dash/pkg/logger"
	"gorm.io/gorm"
)

// TODO: Add database connection and other services as needed
type MiniDashApp struct {
	Database   *gorm.DB
	httpServer *gin.Engine
}

func registerAllRoutes(router *gin.Engine, db *gorm.DB) *gin.Engine {
	// Register project routes
	project_repository := projects.NewProjectRepository(db)
	project_service := projects.NewProjectService(project_repository)
	project_handler := projects.NewProjectHandler(project_service)
	projects.RegisterRoutes(router, project_handler)

	// Register user routes
	user_repository := users.NewUserRepository(db)
	user_service := users.NewUserService(user_repository)
	user_handler := users.NewUserHandler(user_service)
	users.RegisterRoutes(router, user_handler)

	return router
}

func NewMiniDashApp(db *gorm.DB) *MiniDashApp {
	r := gin.New()
	r.Use(logger.GinLogger(logger.MiniLogger()), gin.Recovery()) // Use recovery middleware to handle panics
	r = registerAllRoutes(r, db)

	return &MiniDashApp{
		Database:   db,
		httpServer: r,
	}
}

func (app *MiniDashApp) Run(port int) error {
	// Start the HTTP server
	port_str := fmt.Sprintf(":%d", port)
	return app.httpServer.Run(port_str) // Use the port from config or environment variable
}
