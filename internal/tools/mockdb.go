package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex":{
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason":{
		AuthToken: "456DEF",
		Username: "jason",
	},
	"mario":{
		AuthToken: "789GHI",
		Username: "mario",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex":{
		Coins: 100,
		Username: "alex",
	},
	"jason":{
		Coins: 200,
		Username: "jason",
	},
	"mario":{
		Coins: 300,
		Username: "mario",
	},
}


func(d *mockDB) GetUserLoginDetials(username string) *LoginDetails{
	//simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok{
		return nil
	}

	return &clientData
}

func(d *mockDB) GetUserCoins(username string) *CoinDetails{
	//simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok{
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error{
	return nil
}