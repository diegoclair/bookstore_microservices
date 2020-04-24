package entity

// InitialConfig entity
type InitialConfig struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Schema   string `json:"schema,omitempty"`
}
