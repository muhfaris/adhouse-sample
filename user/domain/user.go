package domain

// User is wrap login user
type User struct {
	Username string `json:"username,omitempty'"`
	Password string `json:"password,omitempty'"`
}

// CreateUser is create new user
func CreateUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
