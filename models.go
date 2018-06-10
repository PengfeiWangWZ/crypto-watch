package main

type OrderBook struct {
	Sequence int `json:"sequence"`
	Bids interface{} `json:"bids"`
	Asks interface{} `json:"asks"`
}

type TickPrice struct {
	Id int `json:"trade_id"`
	Price string `json:"price"`
	Size string `json:"size"`
	Bid string `json:"bid"`
	Ask string `json:"ask"`
	Volume string `json:"volume"`
	Time string `json:"time"`
}

type Option struct {
	Coin string
	Option string
}