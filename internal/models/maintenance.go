package models

// Maintenance JSON 구조체
type Maintenance struct {
	Message []MaintenanceMessage `yaml:"message" json:"message"`
}

type MaintenanceMessage struct {
	Language    string `json:"language,omitempty"` // "name" 대신 "language" 사용
	Title       string `json:"title"`
	Body        string `json:"body"`
	DetailedURL string `json:"detailed_url"`
}
