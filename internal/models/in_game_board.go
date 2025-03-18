package models

// InGameBoard JSON 구조체
type InGameBoard struct {
	Display bool   `mapstructure:"display" json:"display"`
	Url     string `mapstructure:"url" json:"url"`
}
