package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/model"
	"github.com/stephenZ22/mini_dash/internal/service"
)

type CardHandler struct {
	svc service.CardService
}

type CreateCardRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	CardType    uint   `json:"card_type" binding:"required"`
	CreaterID   uint   `json:"creater_id" binding:"required"`
	ProjectId   *uint  `json:"project_id"`
	ParentID    *uint  `json:"parent_id"`
}

func NewCardHandler(svc service.CardService) *CardHandler {
	return &CardHandler{svc: svc}
}

func (h *CardHandler) CreateCard(c *gin.Context) {
	var req CreateCardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	card := &model.Card{
		Name:        req.Name,
		Description: req.Description,
		CardType:    req.CardType,
		ProjectID:   req.ProjectId,
		CreaterID:   req.CreaterID,
		ParentID:    req.ParentID,
	}

	if err := h.svc.CreateCard(card); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "Card created successfully", "data": card})
}
func (h *CardHandler) GetCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	card, err := h.svc.GetCardByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "Card retrieved successfully", "data": card})
}
func (h *CardHandler) UpdateCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	var card model.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	card.ID = uint(id)

	if err := h.svc.UpdateCard(&card); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "Card updated successfully", "data": card})
}
func (h *CardHandler) DeleteCard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	if err := h.svc.DeleteCard(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "Card deleted successfully"})
}
