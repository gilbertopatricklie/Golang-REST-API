package routes

import (
	"example.com/restapi/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)

	
	authenticate := server.Group("/")
	authenticate.Use(middleware.Authenticate)
	authenticate.POST("/events/:id/register", registerEvents)
	authenticate.DELETE("/events/:id/register", cancelRegis)

	
	adminOnly := server.Group("/")
	adminOnly.Use(middleware.Authenticate, middleware.AdminOnly)
	adminOnly.POST("/events", createEvents)
	adminOnly.PUT("/events/:id", updateEvent)
	adminOnly.DELETE("/events/:id", deleteEvent)


	server.POST("/signup", signup)
	server.POST("/login", login)
}