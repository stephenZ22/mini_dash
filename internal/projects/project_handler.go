package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	// TODO add svc
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (ph *ProjectHandler) CreateProject(c *gin.Context) {
	// TODO implement project creation logic
	var req CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0,
		"message": "Project created successfully", "data": gin.H{
			"name":        req.Name,
			"description": req.Description},
	})

}
