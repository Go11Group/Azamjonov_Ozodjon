package middleware

import (
	"fmt"
	"github.com/imtihon/3/Api_gateway/api/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				fmt.Errorf("authorization header is required"))
			return
		}

		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				fmt.Errorf("invalid token: %s", err))
			return
		}

		claims, err := token.ExtractClaims(auth)
		if err != nil || !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				fmt.Errorf("invalid token claims: %s", err))
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
