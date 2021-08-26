package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Serj1c/user-balance/pkg/users"
)

// UserHandler is a handler to work with all things user
type UserHandler struct {
	r *users.Repo
}

// NewUserHandler creates a user handler
func NewUserHandler(r *users.Repo) *UserHandler {
	return &UserHandler{r}
}

// GetBalance handles requests for user's balance
func (uh *UserHandler) GetBalance(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams["user"][0])
	if err != nil {
		http.Error(rw, "Cannot parse user ID", http.StatusBadRequest)
	}
	balance := uh.r.Balance(userID)
	response, err := json.Marshal(balance)
	if err != nil {
		http.Error(rw, "Cannot marshal response", http.StatusInternalServerError)
	}
	rw.Write(response)
}

// Deposit handles requests for depositing money on user's balance
func (uh *UserHandler) Deposit(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams["user"][0])
	if err != nil {
		http.Error(rw, "Cannot parse user ID", http.StatusBadRequest)
	}
	amount, err := strconv.Atoi(queryParams["amount"][0])
	if err != nil {
		http.Error(rw, "Cannot parse amount of money", http.StatusBadRequest)
	}
	err = uh.r.Deposit(userID, amount)
	if err != nil {
		/* .... */
	}
}

// Withdraw handles requests for withdrawl of money out of user's balance
func (uh *UserHandler) Withdraw(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams["user"][0])
	if err != nil {
		http.Error(rw, "Cannot parse user ID", http.StatusBadRequest)
	}
	amount, err := strconv.Atoi(queryParams["amout"][0])
	if err != nil {
		http.Error(rw, "Cannot parse amount of money", http.StatusBadRequest)
	}
	err = uh.r.Withdraw(userID, amount)
	if err != nil {
		/* ... */
	}
}

// Transfer handles requests for transfering money from one user's balance to the other
func (uh *UserHandler) Transfer(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	fromUserID, err := strconv.Atoi(queryParams["from_user"][0])
	if err != nil {
		http.Error(rw, "Cannot parse user ID", http.StatusBadRequest)
	}
	toUserID, err := strconv.Atoi(queryParams["to_user"][0])
	if err != nil {
		http.Error(rw, "Cannot parse user ID", http.StatusBadRequest)
	}
	amount, err := strconv.Atoi(queryParams["amout"][0])
	if err != nil {
		http.Error(rw, "Cannot parse amount of money", http.StatusBadRequest)
	}
	err = uh.r.Transfer(fromUserID, toUserID, amount)
	if err != nil {
		/* ... */
	}
}
