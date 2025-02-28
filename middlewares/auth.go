package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_rest_api/utils"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not Authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not Authorized"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
