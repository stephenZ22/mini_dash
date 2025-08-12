package projects

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/pkg/logger"
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

type UpdateProjectRequest struct {
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
func (ph *ProjectHandler) GetProject(c *gin.Context) {
	idStr := c.Param("id")
	var id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	project, err := ph.svc.GetProject(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "Project retrieved successfully", "data": project})
}

func (ph *ProjectHandler) UpdateProject(c *gin.Context) {

	idStr := c.Param("id")
	var id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}
	var updateProject UpdateProjectRequest
	if err := c.ShouldBindJSON(&updateProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	project := Project{
		Name:        updateProject.Name,
		Description: updateProject.Description,
	}

	if err := ph.svc.UpdateProject(uint(id), &project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "Project updated successfully", "data": project})
}

func (ph *ProjectHandler) DeleteProject(c *gin.Context) {
	idStr := c.Param("id")
	var id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	if err := ph.svc.DeleteProject(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "Project deleted successfully"})
}

func (ph *ProjectHandler) ListProjects(c *gin.Context) {
	pageSize := 10
	pageNum := 1
	if sizeStr := c.Query("page_size"); sizeStr != "" {
		if size, err := strconv.Atoi(sizeStr); err == nil {
			pageSize = size
		}
	}
	if numStr := c.Query("page_num"); numStr != "" {
		if num, err := strconv.Atoi(numStr); err == nil {
			pageNum = num
		}
	}

	logger.MiniLogger().Info("Listing projects", "page_size:", pageSize, "page_num:", pageNum)
	projects, err := ph.svc.ListProjects(pageSize, pageNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "Projects retrieved successfully", "data": projects})
}
