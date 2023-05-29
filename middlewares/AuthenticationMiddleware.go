package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

var jwtKey = []byte("Secret-Key")
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(context echo.Context) error {
	var cred Credentials
	if err := context.Bind(&cred); err != nil {
		return context.String(500, "Binding failed")
	}
	expectedPassword, err := users[cred.Username]

	if !err || expectedPassword != cred.Password {
		return context.NoContent(http.StatusUnauthorized)
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := Claims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err2 := token.SignedString(jwtKey)

	if err2 != nil {
		return context.NoContent(http.StatusInternalServerError)
	}

	context.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return context.String(200, "Login Successful")
}

func Home(context echo.Context) error {
	cookie, err := context.Request().Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			context.NoContent(http.StatusUnauthorized)
		}
		return context.NoContent(http.StatusBadRequest)
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return context.NoContent(http.StatusUnauthorized)
		}
		return context.NoContent(http.StatusBadRequest)
	}

	if !tkn.Valid {
		return context.NoContent(http.StatusUnauthorized)
	}

	return context.String(http.StatusOK, fmt.Sprintf("Hello, %s", claims.Username))
}

func Refresh(context echo.Context) error {
	cookie, err := context.Request().Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			context.NoContent(http.StatusUnauthorized)
		}
		return context.NoContent(http.StatusBadRequest)
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return context.NoContent(http.StatusUnauthorized)
		}
		return context.NoContent(http.StatusBadRequest)
	}
	if !tkn.Valid {
		return context.NoContent(http.StatusUnauthorized)
	}

	//if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	//	return context.NoContent(http.StatusBadRequest)
	//}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err2 := token.SignedString(jwtKey)

	if err2 != nil {
		return context.NoContent(http.StatusInternalServerError)
	}

	context.SetCookie(&http.Cookie{
		Name:    "refresh-token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	return context.NoContent(http.StatusOK)
}
