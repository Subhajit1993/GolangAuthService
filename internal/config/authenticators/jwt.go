package authenticator

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// Function to create a JWT token from user_id
func CreateToken(userId int) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateRefreshToken(userId int, expInHr int) (string, time.Time, error) {
	var err error
	//Creating Refresh Token
	expInHrDuration := time.Duration(expInHr)
	expTime := time.Now().Add(time.Hour * expInHrDuration)
	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = userId
	rtClaims["exp"] = expTime.Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", time.Time{}, err
	}
	return refreshToken, expTime, nil
}

// Function to validate a JWT token

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
