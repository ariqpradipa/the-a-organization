package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bookweb/utils/token"
)

func JWTAuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		err := token.TokenValid(ctx)
		if err != nil{
			ctx.String(http.StatusUnauthorized,"Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
	
}