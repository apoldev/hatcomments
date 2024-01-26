package user

import "github.com/golang-jwt/jwt/v4"

type JwtInfoData struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
}

type JwtData struct {
	jwt.RegisteredClaims
	Info JwtInfoData `json:"info"`
}

type SocialData struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Picture   string `json:"picture"`
}
