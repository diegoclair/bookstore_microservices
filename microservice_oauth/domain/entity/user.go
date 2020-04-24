package entity

// APIUser struct
type APIUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// APIUserLoginRequest struct
type APIUserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
