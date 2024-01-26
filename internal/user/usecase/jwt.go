package usecase

import (
	"cloud_payments/internal/models"
	user2 "cloud_payments/internal/user"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type JwtManager struct {
	HmacSecret    []byte
	Method        jwt.SigningMethod
	ExpireSeconds time.Duration
}

func (m *JwtManager) GetJwtData(tokenStr string) (*user2.JwtData, error) {

	data := &user2.JwtData{}

	token, err := jwt.ParseWithClaims(tokenStr, data, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return m.HmacSecret, nil
	})

	if token.Valid {
		return data, nil
	}

	return nil, err
}

func (m *JwtManager) SignJwt(user *models.User) (*user2.JwtData, string, error) {

	// now := time.Now()

	image := user.Image

	if image == "" {
		name := strings.TrimSpace(user.FirstName + " " + user.LastName)
		image = os.Getenv("APP_URL") + "/user/avatar?name=" + url.PathEscape(name)
	}

	jwtData := user2.JwtData{
		Info: user2.JwtInfoData{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Image:     image,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			//ExpiresAt: &jwt.NumericDate{now.Add(time.Second * m.ExpireSeconds)},
			//IssuedAt:  &jwt.NumericDate{now},
			Subject: strconv.Itoa(int(user.ID)),
		},
	}

	token := jwt.NewWithClaims(m.Method, jwtData)

	tokenString, err := token.SignedString(m.HmacSecret)

	return &jwtData, tokenString, err
}
