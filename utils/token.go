package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, userRole string) (string,error){
	expStr :=  os.Getenv("JWT_EXP_HOURS")
	exp,_ := strconv.Atoi(expStr)


	claims := jwt.MapClaims{
		"user_id" : userID,
		"role" : userRole,
		"exp" : time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	secretKey := os.Getenv("JWT_SECRET")

	signedToken,err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "",err
	}
	return signedToken,nil
}