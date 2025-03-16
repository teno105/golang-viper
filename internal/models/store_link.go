package models

// StoreLink JSON 구조체
type StoreLink struct {
	PlatformType string `json:"platrom_type,omitempty"`
	StoreUrl     string `json:"store_url"`
}
