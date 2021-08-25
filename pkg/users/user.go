package users

// User represents a user
type User struct {
	ID        int32  `json:"ID"`
	Balance   uint64 `json:"balance"`
	Currency  string `json:"currency"`
	CreatedAt string `json:"createdat"`
}
