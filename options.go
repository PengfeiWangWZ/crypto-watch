package main

func getOptions() ([]Option) {
    options := []Option{
        {Coin: "BTC", Option: "BTC-USD"},
        {Coin: "ETH", Option: "ETH-USD"},
        {Coin: "BCH", Option: "BCH-USD"},
    }
    return options
}