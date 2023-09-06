package models

type ServiceNamesIn struct {
	Services []string `json:"service"`
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
