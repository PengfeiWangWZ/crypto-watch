package main


const (
    App = "Crypto Watch"
    refreshFreq = 3
	IconSuccess = "✔"
	IconFail = "✗"
)

func getOptions() ([]Option) {
    options := []Option{
        {Coin: "BTC", Option: "BTC-USD"},
        {Coin: "ETH", Option: "ETH-USD"},
        {Coin: "BCH", Option: "BCH-USD"},
    }
    return options
}