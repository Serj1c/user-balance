package users

// User represents a user
type User struct {
	ID        string  `json:"ID"`
	Balance   float64 `json:"balance"`
	Currency  string  `json:"currency"`
	CreatedAt string  `json:"createdat"`
}
