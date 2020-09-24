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
	tokenString := authHeader[len(bearerSchema):]
	log.Println(tokenString)
	token, err := services.NewJWTService().ValidateToken(tokenString)
	if token == nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		log.Println("id", claims["id"])
		context.Set(common.K_USER_HEADER, claims["id"])
	} else {
		log.Println(err)
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}
