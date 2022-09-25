package jwt

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "pw1",
	"user2": "pw2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Validate(w http.ResponseWriter, r *http.Request) error {
	var tokenStr = ""

	bearer := r.Header.Get("Authorization")
	if bearer != "" {
		token := strings.Split(bearer, " ")
		tokenStr = token[1]
	} else {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return fmt.Errorf("Unauthorized")
			}
			w.WriteHeader(http.StatusBadRequest)
			return fmt.Errorf("Bad Request")
		}
		if cookie != nil {
			tokenStr = cookie.Value
		}

		if bearer == "" && cookie == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return fmt.Errorf("Unauthorized")
		}
	}

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return fmt.Errorf("Unauthorized")
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return fmt.Errorf("Unauthorized")
	}

	return nil
}
