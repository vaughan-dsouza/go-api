package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/vaughan-dsouza/go-api/api"
	"github.com/vaughan-dsouza/go-api/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request){
	var params = api.CoinbalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.internalErrorHandler{w}
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinbalancResponse{
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}