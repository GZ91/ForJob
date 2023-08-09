package models

import "github.com/golang-jwt/jwt/v4"

type ServiceStr struct {
	Service string `json:"service"`
}

type StructURL struct {
	ID          string `json:"uuid" db:"uuid"`
	ShortURL    string `json:"short_url" db:"ShortURL"`
	OriginalURL string `json:"original_url" db:"OriginalURL"`
	UserID      string `json:"user_id"`
	DeletedFlag bool
}

type RequestData struct {
	URL string `json:"longLink"`
}

type ReturnData struct {
	LongLink  string `json:"longLink"`
	ShortLink string `json:"shortLink"`
}

type ResultReturn struct {
	Result string `json:"result"`
}

type ReturnedStructURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type CtxString string

type StructDelURLs struct {
	URL    string
	UserID string
}

type Claims struct {
	jwt.RegisteredClaims
	Token string
}
