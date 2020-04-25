package entity

import (
	"fmt"
	"strings"

	"github.com/diegoclair/go_utils-lib/resterrors"
)

// User struct
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// Validate to validate a user data
func (user *User) Validate() *resterrors.RestErr {

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return resterrors.NewBadRequestError("Invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	err := user.validadePassword()
	if err != nil {
		return err
	}

	return nil
}

func (user *User) validadePassword() *resterrors.RestErr {
	fmt.Println("cheguei aqui", user.Password)
	if user.Password == "" || len(user.Password) < 8 {
		return resterrors.NewBadRequestError("Password need at least 8 caracters")
	}

	return nil
}

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

//LoginRequest struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
