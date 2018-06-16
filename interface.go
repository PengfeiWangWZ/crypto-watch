package main

import (
	"fmt"
	"time"
	"sort"
	"github.com/marcusolsson/tui-go"
)

var (
	updateChan = make(chan *ProductInfo)
)

func userInterface(product string) {
	grid := tui.NewGrid(0, 0)

	coinLabel := tui.NewLabel(fmt.Sprintf("Coin: %s", product))
	priceLabel := tui.NewLabel("Price: ")
	grid.AppendRow(tui.NewHBox(coinLabel, priceLabel))

	headerBar := tui.NewVBox(grid)
	headerBar.SetTitle(App)
	headerBar.SetBorder(true)

	asksAdds := tui.NewVBox()
	asksAdds.SetSizePolicy(tui.Expanding, tui.Maximum)
	asksBox := tui.NewVBox(tui.NewVBox(
		tui.NewHBox(tui.NewLabel("Price"), tui.NewLabel("Volume"))),
		asksAdds)
	asksBox.SetTitle("Asks")
	asksBox.SetBorder(true)

	bidsAdds := tui.NewVBox()
	bidsAdds.SetSizePolicy(tui.Expanding, tui.Maximum)
	bidsBox := tui.NewVBox(tui.NewVBox(
		tui.NewHBox(tui.NewLabel("Price"), tui.NewLabel("Volume"))),
		bidsAdds)
	bidsBox.SetTitle("Bids")
	bidsBox.SetBorder(true)

	orderBookBox := tui.NewHBox(asksBox, bidsBox)
	orderBookBox.SetSizePolicy(tui.Preferred, tui.Expanding)
	root := tui.NewVBox(headerBar, orderBookBox)
	ui := tui.New(root)
	
	ui.SetKeybinding("Ctrl+C", func() { ui.Quit() })
	ui.SetKeybinding("ESC", func() { ui.Quit() })
	ui.SetKeybinding("Q", func() { ui.Quit() })

	go update(ui, priceLabel, asksAdds, bidsAdds)
	go fetch(product)
	
	err := ui.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func update(ui tui.UI, priceLabel *tui.Label, asksAdds, bidsAdds *tui.Box) {
	defer close(updateChan)
	for {
		select {
		case data := <-updateChan:
			tickPrice, orderBook := data.TickPriceInfo, data.OrderBookInfo
			ui.Update(func() {
				if tickPrice != nil {
					priceLabel.SetText(fmt.Sprintf("Price: %s", tickPrice.Price))
				}
				if orderBook != nil {
					asks := getDetailedOrders(*orderBook, false)
					bids := getDetailedOrders(*orderBook, true)
					sort.Sort(sort.Reverse(ByPrice(asks)))
					sort.Sort(ByPrice(bids))
					for index := 0; index < len(asks); index ++ {
						asksAdds.Remove(index)
						bidsAdds.Remove(index)
					}
					for _, order := range asks {
						asksAdds.Prepend(tui.NewHBox(
							tui.NewLabel(order.Price),
							tui.NewLabel(order.Volume),
						))
					}
					for _, order := range bids {
						bidsAdds.Prepend(tui.NewHBox(
							tui.NewLabel(order.Price),
							tui.NewLabel(order.Volume),
						))
					}
				}
			})
		}
	}
}

func fetch(product string) {
	for {
		orderBook, err := getOrderBookByProduct(product)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		tickPrice, err := getLatestPriceByProduct(product)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		updateChan <- &ProductInfo{
			TickPriceInfo: tickPrice,
			OrderBookInfo: orderBook,
		}
		time.Sleep(refreshFreq * time.Second)
	}
}