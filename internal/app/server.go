package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/handler"
	"github.com/stephenZ22/mini_dash/internal/repository"
	"github.com/stephenZ22/mini_dash/internal/router"
	"github.com/stephenZ22/mini_dash/internal/service"
	"github.com/stephenZ22/mini_dash/pkg/logger"
	"gorm.io/gorm"
)

// TODO: Add database connection and other services as needed
type MiniDashApp struct {
	Database   *gorm.DB
	httpServer *gin.Engine
}

func registerAllRoutes(engin *gin.Engine, db *gorm.DB) *gin.Engine {
	// Register project routes
	project_repository := repository.NewProjectRepository(db)
	project_service := service.NewProjectService(project_repository)
	project_handler := handler.NewProjectHandler(project_service)
	router.RegisterProjectRoutes(engin, project_handler)

	// Register user routes
	user_repository := repository.NewUserRepository(db)
	user_service := service.NewUserService(user_repository)
	user_handler := handler.NewUserHandler(user_service)
	router.RegisterUserRoutes(engin, user_handler)

	card_repository := repository.NewCardRepository(db)
	card_service := service.NewCardService(card_repository)
	card_handler := handler.NewCardHandler(card_service)
	router.RegisterCardRoutes(engin, card_handler)

	return engin
}

func NewMiniDashApp(db *gorm.DB) *MiniDashApp {
	engin := gin.New()
	engin.Use(logger.GinLogger(logger.MiniLogger()), gin.Recovery()) // Use recovery middleware to handle panics
	engin = registerAllRoutes(engin, db)

	return &MiniDashApp{
		Database:   db,
		httpServer: engin,
	}
}

func (app *MiniDashApp) Run(port int) error {
	// Start the HTTP server
	port_str := fmt.Sprintf(":%d", port)
	return app.httpServer.Run(port_str) // Use the port from config or environment variable
}
