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
func (r *Repo) Deposit(userID string, amount float64) error {
	user := &User{}
	err := r.db.QueryRow("SELECT balance from users WHERE id = $1", userID).Scan(&user.Balance)
	if err == sql.ErrNoRows {
		tx, err := r.db.Begin()
		if err != nil {
			return ErrDBQuery
		}
		defer tx.Rollback()
		_, err = tx.Exec("INSERT into users(id, balance) VALUES($1, $2)", userID, amount)
		if err != nil {
			return ErrDBQuery
		}
		_, err = tx.Exec("INSERT into deposits(to_user_id, amount) VALUES($1, $2)", userID, amount)
		if err != nil {
			return ErrDBQuery
		}
		err = tx.Commit()
		if err != nil {
			return ErrDBQuery
		}
		return nil
	}
	tx, err := r.db.Begin()
	if err != nil {
		return ErrDBQuery
	}
	defer tx.Rollback()
	newBalance := user.Balance + amount
	_, err = tx.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalance, userID)
	if err != nil {
		return ErrDBQuery
	}
	_, err = tx.Exec("INSERT into deposits(to_user_id, amount) VALUES($1, $2)", userID, amount)
	if err != nil {
		return ErrDBQuery
	}
	err = tx.Commit()
	if err != nil {
		return ErrDBQuery
	}
	return nil
}

// Withdraw takes money out of user's balance
func (r *Repo) Withdraw(userID string, amount float64) error {
	user := &User{}
	tx, err := r.db.Begin()
	if err != nil {
		return ErrDBQuery
	}
	defer tx.Rollback()
	err = tx.QueryRow("SELECT balance from users WHERE id = $1", userID).Scan(&user.Balance)
	if err == sql.ErrNoRows {
		return ErrNoUser
	}
	if user.Balance >= amount {
		newBalance := user.Balance - amount
		_, err = tx.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalance, userID)
		if err != nil {
			return ErrDBQuery
		}
		_, err = tx.Exec("INSERT into withdrawals(from_user_id, amount) VALUES($1, $2)", userID, amount)
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

// Transfer transfers money from one user's balance to the other
func (r *Repo) Transfer(fromUserID, toUserID string, amount float64) error {
	fromUser := &User{}
	toUser := &User{}
	tx, err := r.db.Begin()
	if err != nil {
		return ErrDBQuery
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
		_, err = tx.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalanceFrom, fromUserID)
		if err != nil {
			return ErrDBQuery
		}
		_, err = tx.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalanceTo, toUserID)
		if err != nil {
			return ErrDBQuery
		}
		_, err = tx.Exec("INSERT into transactions(from_user_id, to_user_id, amount) VALUES($1, $2, $3)", fromUserID, toUserID, amount)
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
func (r *Repo) Balance(userID string) (float64, error) {
	user := &User{}
	err := r.db.QueryRow("SELECT balance FROM users WHERE id = $1", userID).Scan(&user.Balance)
	if err == sql.ErrNoRows {
		return -1, ErrNoUser
	}
	return user.Balance, nil
}

// UserBalanceOperation represents an operation with user's balance such as deposit of money or their withdrawal
type UserBalanceOperation struct {
	ID         int     `json:"id"`
	FromUserID string  `json:"from_user_id"`
	ToUserID   string  `json:"to_user_id"`
	Amount     float64 `json:"amount"`
	CreatedAt  string  `json:"created_at"`
	Comment    string  `json:"comment"`
}

// List returns all user's operations with the balance
func (r *Repo) List(userID, sortBy, sortOrder string, perPage, offset int) ([]*UserBalanceOperation, error) {
	operations := make([]*UserBalanceOperation, 0, 10)
	sorting := fmt.Sprintf(" ORDER BY %s %s", sortBy, sortOrder)
	pagination := fmt.Sprintf(" LIMIT %d OFFSET %d", perPage, offset)
	rows, err := r.db.Query(`SELECT id, from_user_id, to_user_id, amount, created_at, comment FROM deposits 
	WHERE to_user_id = $1
	UNION ALL SELECT id, from_user_id, to_user_id, amount, created_at, comment FROM withdrawals 
	WHERE from_user_id = $1
	UNION ALL SELECT id, from_user_id, to_user_id, amount, created_at, comment FROM transactions 
	WHERE from_user_id = $1
	UNION ALL SELECT id, from_user_id, to_user_id, amount, created_at, comment FROM transactions 
	WHERE to_user_id = $1`+sorting+pagination, userID)
	if err != nil {
		fmt.Println(err)
		return nil, ErrDBQuery
	}
	defer rows.Close()
	for rows.Next() {
		item := &UserBalanceOperation{}
		err := rows.Scan(&item.ID, &item.FromUserID, &item.ToUserID, &item.Amount, &item.CreatedAt, &item.Comment)
		if err != nil {
			return nil, ErrDBQuery
		}
		operations = append(operations, item)
	}
	return operations, nil
}
