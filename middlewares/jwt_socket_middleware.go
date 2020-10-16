package middlewares

import (
	"errors"
	"github.com/ElegantSoft/shabahy/common"
	"github.com/ElegantSoft/shabahy/services"
	"github.com/dgrijalva/jwt-go"
	socketio "github.com/googollee/go-socket.io"
	"net/url"
)

func AuthorizeSocket(s socketio.Conn) (error, uint) {
	rawQuery := s.URL().RawQuery
	query, _ := url.ParseQuery(rawQuery)
	tokenSlice := query["access_token"]
	var tokenString string
	if len(tokenSlice) > 0 {
		tokenString = tokenSlice[0]
	} else {
		tokenString = ""
		return errors.New("no token"), 0
	}

	token, err := services.NewJWTService().ValidateToken(tokenString)
	if token == nil {
		return errors.New("not valid token"), 0
	}
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		var id = claims["id"]
		return nil, common.GetIdFromCtx(id)
	} else {
		return err, 0
	}
}
