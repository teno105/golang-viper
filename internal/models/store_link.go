package models

// StoreLink JSON 구조체
type StoreLink struct {
	PlatformType string `json:"platform_type,omitempty"`
	StoreUrl     string `json:"store_url"`
}
