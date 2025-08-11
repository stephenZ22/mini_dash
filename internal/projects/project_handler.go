package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	svc *ProjectService
}

func NewProjectHandler(svc *ProjectService) *ProjectHandler {
	return &ProjectHandler{
		svc: svc,
	}
}

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (ph *ProjectHandler) CreateProject(c *gin.Context) {
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := ph.svc.CreateProject(req.Name, req.Description); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1,
		"message": "Project created successfully", "data": gin.H{
			"name":        req.Name,
			"description": req.Description},
	})
}
