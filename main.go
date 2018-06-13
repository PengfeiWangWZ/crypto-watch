package main

import (
	"fmt"
)

func main() {
	coins := getOptions()
	product, err := selectCoin(coins)
	if err != nil {
		fmt.Printf("%s %s\n", IconFail, "Failed to select coin")
		return
	}
	userInterface(coins[product].Option)
}