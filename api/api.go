package api

import (
	"encoding/json"
	"net/http"
)

// Coin balanceParams struct
type CoinbalanceParams struct {
	Username string
}

// Coin balanceParams Response
type CoinbalancResponse struct {
	//Success Code 200
	Code int

	//Account Balance
	Balance int64
}

// Error Response
type Error struct {
	//Error Code 400
	Code int

	//Account Balance
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpecter Error Occured.", http.StatusInternalServerError)
	}
)
