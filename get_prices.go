package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func getOrderBookByProduct(product string) (*OrderBook, error) {
	baseUrl := "https://api.gdax.com/products/%s/book?level=2"
	url := fmt.Sprintf(baseUrl, product)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	productOrderBook := &OrderBook{}
	err = json.Unmarshal(response, productOrderBook)
	return productOrderBook, err
}

func getLatestPriceByProduct(product string) (*TickPrice, error) {
	baseUrl := "https://api.gdax.com/products/%s/ticker"
	url := fmt.Sprintf(baseUrl, product)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	productTicker := &TickPrice{}
	err = json.Unmarshal(response, productTicker)
	return productTicker, err
}

func getDetailedOrders(orderBook OrderBook, isBid bool) ([]Order) {
	orders := []Order{}
	if isBid {
		for _, bid := range orderBook.Bids {
			orders = append(orders, Order{
				Price:  bid[0].(string),
				Volume: bid[1].(string),
			})
		}
	} else {
		for _, bid := range orderBook.Bids {
			orders = append(orders, Order{
				Price:  bid[0].(string),
				Volume: bid[1].(string),
			})
		}
	}
	return orders
}