package middleware

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func (c *casbinPermission) GetRole(ctx *gin.Context) (string, int) {
	role := ctx.GetHeader("Authorization")

	if role == "" {
		return "unauthorization", http.StatusUnauthorized
	}

	return role, 0
}

func (c *casbinPermission) CheckPermission(ctx *gin.Context) (bool, error) {

	act := ctx.Request.Method
	sub, status := c.GetRole(ctx)
	obj := ctx.Request.URL.Path

	if status != 0 {
		return false, fmt.Errorf("error geting role from Authorization")
	}

	allow, err := c.enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, fmt.Errorf("error with checking permission: %s", err)

	}

	return allow, nil
}

func CheckPermissionMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	casbinHandler := casbinPermission{
		enforcer: enforcer,
	}

	return func(ctx *gin.Context) {
		result, err := casbinHandler.CheckPermission(ctx)

		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
		}

		if !result {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "StatusUnauthorized",
				"message": "Unauthorized",
			})
		}

		ctx.Next()
	}
}
