package function

import (
	"ehsan_esmaeili/config"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)

	claims["iat"] = time.Now().Unix() //create time
	// claims["uid"] = userId
	// claims["agent"] = agent
	// claims["group"] = config.Group_Clent
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix() //expire time

	tokenString, err := token.SignedString([]byte(config.Jwt_Signing_Key))

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil

}
