package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var mySigningKey1 = []byte("mySecretKey")

type MyClaims1 struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成JWT令牌
func generateJWT1(userID string, username string) (string, error) {
	claims := MyClaims1{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Lorry", //发行者
			Subject:   "test",  //主题
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	//创建新的JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//使用密钥签名JWT
	signedToken, err := token.SignedString(mySigningKey1)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//验证JWT令牌

func validateJWT1(signedToken string) (*MyClaims1, error) {
	//解析JWT并验证
	token, err := jwt.ParseWithClaims(signedToken, &MyClaims1{}, func(token *jwt.Token) (interface{}, error) {
		//这里只验证签名的密钥
		return mySigningKey1, nil
	})
	if err != nil {
		return nil, err
	}
	//验证有效性
	if claims, ok := token.Claims.(*MyClaims1); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

// Middleware:验证JWT
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从请求头中获取Bearer Token
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 7 || authHeader[:7] != "Bearer" {
			c.JSON(401, gin.H{
				"error": "Authorization header missing or invalid"})
			c.Abort()
			return
		}
		//提取Token部分
		tokenString := authHeader[7:]

		//验证Token
		Claims, err := validateJWT1(tokenString)
		if err != nil {
			c.JSON(401, gin.H{})
			c.Abort()
			return
		}
		c.Set("UserID", Claims.UserID)
		c.Set("username", Claims.Username)
		c.Next()
	}
}
func main() {
	r := gin.Default()

	r.POST("/test", func(c *gin.Context) {
		var loginData struct {
			UserID   string `json:"user_id"`
			Username string `json:"username"`
		}
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//生成JWT
		tokenString, err := generateJWT1(loginData.UserID, loginData.Username)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		//返回生成的JWT
		c.JSON(200, gin.H{"token": tokenString})
	})

	//需要认证的路由
	r.GET("/test1", JWTMiddleware(), func(c *gin.Context) {
		//获取之前存储的用户ID
		userID, _ := c.Get("UserID")
		username := c.MustGet("username").(string)
		//返回用户ID
		c.JSON(200, gin.H{"message": "Hello, World!", "userId": userID, "username": username})
	})

	r.Run(":8080")
}
