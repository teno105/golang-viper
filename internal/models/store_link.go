package models

// StoreLink JSON 구조체
type StoreLink struct {
	PlatformType string `mapstructure:"platform_type" json:"platform_type"`
	StoreUrl     string `mapstructure:"store_url" json:"store_url"`
}
