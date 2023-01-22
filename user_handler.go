package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHTTPHandler struct {
	service *UserService
	r       *gin.Engine
}

func RegisterUserHTTPHandler(r *gin.Engine, s *UserService) {
	// Create a new HTTP handler for the user service
	h := &UserHTTPHandler{r: r, service: s}

	// Define routes for the handler
	// r := gin.Default()
	h.r.POST("/users", h.Create)
	h.r.GET("/users/:id", h.Get)
	// r.PUT("/users/:id", h.Update)
	// r.DELETE("/users/:id", h.Delete)

	// return r
}

func (h *UserHTTPHandler) Get(c *gin.Context) {
	// Get the user ID from the request parameters
	userID := c.Param("id")

	id, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the service to retrieve the user from the database
	user, err := h.service.FindUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, NewResponse("user retrieved", user))
}

func (h *UserHTTPHandler) Create(c *gin.Context) {
	// Get user data from the request body
	var u AddUserRequestBody
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the service to create a new user in the database
	user := User{Name: u.Name}
	newUser, err := h.service.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, NewResponse("user created", newUser))
}
