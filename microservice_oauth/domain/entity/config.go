package entity

// InitialConfig entity
type InitialConfig struct {
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Host      string `json:"host,omitempty"`
	Port      string `json:"port,omitempty"`
	DBName    string `json:"db_name,omitempty"`
	DBDefault string `json:"db_default,omitempty"`
}
