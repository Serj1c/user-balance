package users

import (
	"database/sql"
	"fmt"
)

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

var (
	// ErrNoUser is raised when user does not exist in db
	ErrNoUser = fmt.Errorf("User does not exist")
	// ErrNotEnoughMoney is raised when user exists but does not have enough money
	ErrNotEnoughMoney = fmt.Errorf("User does not have enough money")
	// ErrDBQuery is raised when there is an error during communication with db
	ErrDBQuery = fmt.Errorf("DB error")
)

// Deposit adds money to a user's balance
func (r *Repo) Deposit(userID string, amount int) error {
	user := &User{}
	err := r.db.QueryRow("SELECT balance from users WHERE id = $1", userID).Scan(&user.Balance)
	if err == sql.ErrNoRows {
		_, err := r.db.Exec("INSERT into users(id, balance) VALUES($1, $2)", userID, amount)
		if err != nil {
			return ErrDBQuery
		}
		_, err = r.db.Exec("INSERT into deposits(user_id, amount) VALUES($1, $2)", userID, amount)
		if err != nil {
			return ErrDBQuery
		}
		return nil
	}
	newBalance := user.Balance + amount
	_, err = r.db.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalance, userID)
	if err != nil {
		return ErrDBQuery
	}
	_, err = r.db.Exec("INSERT into deposits(user_id, amount) VALUES($1, $2)", userID, amount)
	if err != nil {
		return ErrDBQuery
	}
	return nil
}

// Withdraw takes money out of user's balance
func (r *Repo) Withdraw(userID string, amount int) error {
	user := &User{}
	err := r.db.QueryRow("SELECT balance from users WHERE id = $1", userID).Scan(&user.Balance)
	if err == sql.ErrNoRows {
		return ErrNoUser
	}
	if user.Balance >= amount {
		newBalance := user.Balance - amount
		_, err = r.db.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalance, userID)
		if err != nil {
			return fmt.Errorf("Money withdrawal has failed")
		}
		_, err = r.db.Exec("INSERT into withdrawls(user_id, amount) VALUES($1, $2)", userID, amount)
		if err != nil {
			return fmt.Errorf("Money withdrawal has failed")
		}
		return nil
	}
	return ErrNotEnoughMoney
}

// Transfer transfers money from one user's balance to the other
func (r *Repo) Transfer(fromUserID, toUserID string, amount int) error {
	fromUser := &User{}
	toUser := &User{}
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("Money transfer has failed at the very beginning")
	}
	defer tx.Rollback()
	err = tx.QueryRow("SELECT balance from users WHERE id = $1", fromUserID).Scan(&fromUser.Balance)
	if err == sql.ErrNoRows {
		return ErrNoUser
	}
	err = tx.QueryRow("SELECT balance from users WHERE id = $1", toUserID).Scan(&toUser.Balance)
	if err == sql.ErrNoRows {
		return ErrNoUser
	}
	if fromUser.Balance >= amount {
		newBalanceFrom := fromUser.Balance - amount
		newBalanceTo := toUser.Balance + amount
		_, err = r.db.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalanceFrom, fromUserID)
		if err != nil {
			return ErrDBQuery
		}
		_, err = r.db.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalanceTo, toUserID)
		if err != nil {
			return ErrDBQuery
		}
		_, err = r.db.Exec("INSERT into transactions(from_user_id, to_user_id, amount) VALUES($1, $2, $3)", fromUserID, toUserID, amount)
		if err != nil {
			return ErrDBQuery
		}
		err = tx.Commit()
		if err != nil {
			return ErrDBQuery
		}
		return nil
	}
	return ErrNotEnoughMoney
}

// Balance returns user's balance
func (r *Repo) Balance(userID string) (int, error) {
	user := &User{}
	err := r.db.QueryRow("SELECT balance FROM users WHERE id = $1", userID).Scan(&user.Balance)
	if err == sql.ErrNoRows {
		return -1, ErrNoUser
	}
	return user.Balance, nil
}
