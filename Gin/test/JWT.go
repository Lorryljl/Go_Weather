package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var mySigningKey = []byte("mysecretkey")

type MyClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func generateJWT(userID string, username string) (string, error) {
	//创建JWT的Claims(载荷)
	claims := MyClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MyTest",
			Subject:   "MyTesting",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	//创建新的JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//使用密钥签名JWT
	signedToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil

}

// 验证JWT令牌
func validateJwt(tokenString string) (*MyClaims, error) {
	//解析JWT并验证
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	//验证有效性
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
func main() {
	tokenString, err := generateJWT("123456", "lorry")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("token值为:", tokenString)
	clams, err := validateJwt(tokenString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(clams.UserID)
	fmt.Println(clams.Username)
}
