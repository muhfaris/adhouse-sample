package structures

// LoginRead is wrap login user
type LoginRead struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
