package store

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStore interface {
	CreateUser(email string, password string) error
	GetUser(email string) (*User, error)
}
