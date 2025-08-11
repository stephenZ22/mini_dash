package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/projects"
	"gorm.io/gorm"
)

// TODO: Add database connection and other services as needed
type MiniDashApp struct {
	Database   *gorm.DB
	httpServer *gin.Engine
}

func NewMiniDashApp(db *gorm.DB) *MiniDashApp {
	r := gin.Default()
	project_repository := projects.NewProjectRepository(db)
	project_service := projects.NewProjectService(project_repository)
	prject_handler := projects.NewProjectHandler(project_service)
	projects.RegisterRoutes(r, prject_handler)
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
