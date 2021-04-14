package middleware

import (
	"ehsan_esmaeili/config"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func CheckAuthentication(w http.ResponseWriter, r *http.Request) error {
	token, err := jwt.Parse(r.Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(config.Jwt_Signing_Key), nil
	})
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("err")
		return err
	}

	if !token.Valid {
		fmt.Println("is not valid")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("error")
		return fmt.Errorf("is not valid")
	}
	return nil
}
