package viewmodel

// PublicUser - return a struct with few fields
type PublicUser struct {
	ID        int64  `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// PrivateUser - return a struct with all data. It's need a token to request
type PrivateUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
