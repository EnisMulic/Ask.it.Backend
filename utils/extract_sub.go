package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ExtractSubFromJwt(r *http.Request) (string, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims) 
	if ok && token.Valid { 
		sub, _ := claims["sub"].(string)

		return sub, nil
	}

	return "", err
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		aud := os.Getenv("JWT_AUDIENCE")
		checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)

		if !checkAudience {
			return nil, fmt.Errorf(("invalid aud"))
		}
		

		iss := os.Getenv("JWT_ISSUER")
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return nil, fmt.Errorf(("invalid iss"))
		}

		chechExp := token.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().UTC().Unix(), true)
		if !chechExp {
			return nil, fmt.Errorf(("token has expired"))
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}