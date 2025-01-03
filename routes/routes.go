package routes

import (
	"example.com/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	// Create a new group for authenticated routes
	authenticated_group := server.Group("/")
	authenticated_group.Use(middlewares.Authenticate)
	authenticated_group.PUT("/events/:id", updateEvent)
	authenticated_group.DELETE("/events/:id", deleteEvent)
	authenticated_group.POST("/events/:id/register", registerForEvent)
	authenticated_group.DELETE("/events/:id/register", cancelRegistration)
	// Same group of authenticated routes, but with a different middleware implementation
	server.POST("/events", middlewares.Authenticate, createEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
