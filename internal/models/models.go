package models

type ServiceNamesIn struct {
	Services []string `json:"service"`
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

type ReturnedStructURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type CtxString string

type StructDelURLs struct {
	URL    string
	UserID string
}
