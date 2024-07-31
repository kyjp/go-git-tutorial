package middlewares

import (
	"log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"web/app/services"
)

func AuthMiddleware(authService services.IAuthService) gin.HandlerFunc {
	return 	func (ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			log.Println("hogehogehogehoge")
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			log.Println("fugafugafugafuga")
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")
		user, err := authService.GetUserFromToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			log.Println("piyopiyopiyopiyo")
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}