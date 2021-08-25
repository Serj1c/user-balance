package users

import "database/sql"

// Repo represents a database
type Repo struct {
	db *sql.DB
}

// NewRepo returns an instance of a Repo
func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// Deposit adds money to a user's balance
func (r *Repo) Deposit(userID, amount int) error {
	return nil
}

// Withdraw takes money out of user's balance
func (r *Repo) Withdraw(userID, amount int) error {
	return nil
}

// Transfer transfers money from one user's balance to the other
func (r *Repo) Transfer(fromUserID, toUserID int, amount int) error {
	return nil
}

// Balance returns user's balance
func (r *Repo) Balance(userID int) int {
	return userID*100 + 1
}
