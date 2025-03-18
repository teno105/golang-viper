package models

// Maintenance JSON 구조체
type Maintenance struct {
	Message []MaintenanceMessage `mapstructure:"message" json:"message"`
}

type MaintenanceMessage struct {
	Language    string `mapstructure:"language" json:"language,omitempty"`
	Title       string `mapstructure:"title" json:"title"`
	Body        string `mapstructure:"body" json:"body"`
	DetailedURL string `mapstructure:"detailed_url" json:"detailed_url"`
}
