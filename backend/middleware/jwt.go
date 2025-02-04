package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct{
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}
var jwtKey []byte
func init(){
	err := godotenv.Load(".env")
	if err != nil {
		panic("error load .env")
	}
	secretJwt := os.Getenv("JWT_KEY")
	jwtKey = []byte(secretJwt)
}

func GenerateJWT(username, role string)(string, error){
	waktuKadaluarsa := time.Now().Add(30*time.Minute)
	claims := &Claims{
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(waktuKadaluarsa),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}