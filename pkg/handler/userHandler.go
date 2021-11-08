package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Serj1c/user-balance/pkg/users"
	"github.com/Serj1c/user-balance/pkg/util"
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
	userID := queryParams["user"][0]
	balance, err := uh.r.Balance(userID)
	if err == users.ErrNoUser {
		http.Error(rw, "User does not exist", http.StatusBadRequest)
	} else {
		if _, ok := queryParams["currency"]; ok {
			currency := queryParams["currency"][0]
			balance = balance * excangeRateAPIcall(currency)
		}
		response, err := json.Marshal(balance)
		if err != nil {
			http.Error(rw, "Cannot marshal response", http.StatusInternalServerError)
		}
		rw.Write(response)
	}
}

// Deposit handles requests for depositing money on user's balance
func (uh *UserHandler) Deposit(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userID := queryParams["user"][0]
	amount, err := strconv.ParseFloat(queryParams["amount"][0], 64)
	if err != nil {
		http.Error(rw, "Cannot parse amount of money", http.StatusBadRequest)
	}
	if amount > 0 {
		err = uh.r.Deposit(userID, amount)
		switch err {
		case nil:
		case users.ErrDBQuery:
			http.Error(rw, "Server fault", http.StatusInternalServerError)
		}
	} else {
		http.Error(rw, "Deposit of only positive sums is allowed", http.StatusBadRequest)
	}
}

// Withdraw handles requests for withdrawl of money out of user's balance
func (uh *UserHandler) Withdraw(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userID := queryParams["user"][0]
	amount, err := strconv.ParseFloat(queryParams["amount"][0], 64)
	if err != nil {
		http.Error(rw, "Cannot parse amount of money", http.StatusBadRequest)
	}
	if amount > 0 {
		err = uh.r.Withdraw(userID, amount)
		if err == users.ErrNoUser {
			http.Error(rw, "User does not exist", http.StatusBadRequest)
		} else if err == users.ErrNotEnoughMoney {
			http.Error(rw, "User does not have enough money", http.StatusBadRequest)
		}
	} else {
		http.Error(rw, "Withdrawal of only positive sums is allowed", http.StatusBadRequest)
	}
}

// Transfer handles requests for transfering money from one user's balance to the other
func (uh *UserHandler) Transfer(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	fromUserID := queryParams["from_user"][0]
	toUserID := queryParams["to_user"][0]
	amount, err := strconv.ParseFloat(queryParams["amount"][0], 64)
	if err != nil {
		http.Error(rw, "Cannot parse amount of money", http.StatusBadRequest)
	}
	if amount > 0 {
		err = uh.r.Transfer(fromUserID, toUserID, amount)
		switch err {
		case nil:
		case users.ErrNotEnoughMoney:
			http.Error(rw, "User does not have enough money", http.StatusBadRequest)
		case users.ErrNoUser:
			http.Error(rw, "One or both users do not exist", http.StatusBadRequest)
		case users.ErrDBQuery:
			http.Error(rw, "Internal error", http.StatusInternalServerError)
		}
	} else {
		http.Error(rw, "Transfer of only positive sums is allowed", http.StatusBadRequest)
	}
}

// ListAllOperations lists all operations performed on user's balance
func (uh *UserHandler) ListAllOperations(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userID := queryParams["user"][0]
	operations, err := uh.r.List(userID)
	if err != nil {
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
	err = util.ToJSON(operations, rw)
	if err != nil {
		http.Error(rw, "Internal error", http.StatusInternalServerError)
	}
}

func dateTransformer() string {
	t := time.Now().Local().AddDate(0, 0, -1)
	s := t.Format("2006-01-02")
	return s
}

func excangeRateAPIcall(currency string) float64 {
	date := dateTransformer()
	resp, err := http.Get("http://api.exchangeratesapi.io/v1/" + date + "?access_key=36d585d941651b79dd7d412d57dc66ff&base=EUR&symbols=RUB," + currency)
	if err != nil {
		return -1
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	result := make(map[string]interface{})
	json.Unmarshal([]byte(data), &result)
	rates := result["rates"].(map[string]interface{})
	curString := fmt.Sprint(rates[currency])
	rubString := fmt.Sprint(rates["RUB"])
	cur, _ := strconv.ParseFloat(curString, 64)
	rub, _ := strconv.ParseFloat(rubString, 64)
	excangeRate := cur / rub
	return excangeRate
}
