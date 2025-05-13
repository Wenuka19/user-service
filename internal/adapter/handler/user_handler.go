package handler

import (
	"github.com/Wenuka19/user-service/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserHandler struct {
	service domain.UserRepository
}

func NewUserHandler(service domain.UserRepository) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/", h.CreateOrUpdateUser)
	r.GET("/:id", h.GetUser)
}

func (h *UserHandler) CreateOrUpdateUser(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user.UpdatedAt = time.Now()
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	if err := h.service.Save(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
