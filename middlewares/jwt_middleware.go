package middlewares

import (
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthorizeJWT(context *gin.Context) {
	const bearerSchema = "Bearer "
	authHeader := context.GetHeader("Authorization")
	if len(authHeader) < 10 {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString := authHeader[len(bearerSchema):]
	log.Println(tokenString)
	token, err := services.NewJWTService().ValidateToken(tokenString)
	if token == nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"]
		context.Set(common.KUserHeader, id)
		context.Next()
	} else {
		log.Println(err)
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
