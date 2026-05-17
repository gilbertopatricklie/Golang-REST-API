package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
 
func AdminOnly(context *gin.Context) {
	role := context.GetString("role")
	if role != "admin" {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Admin only"})
		return
	}
	context.Next()
}