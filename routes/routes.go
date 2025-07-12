package routes

import (
	"example.com/restapi/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//API ROUTES
	server.GET("/events", getEvents)
	server.GET("/events/:id", getSingleEvent)

	authenticate := server.Group("/")
	authenticate.Use(middleware.Authenticate)
	authenticate.POST("/events", createEvents)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events:id", deleteEvent)
	authenticate.POST("/events/:id/register", registerEvents)
	authenticate.DELETE("/events/:id/register", cancelRegis)

	//USERS ROUTE
	server.POST("/signup", signup)
	server.POST("/login", login)
}