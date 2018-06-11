package main

import (
	"fmt"
	"github.com/marcusolsson/tui-go"
)

func userInterface(product string, tickPrice *TickPrice, asks []Order, bids []Order) {
	grid := tui.NewGrid(0, 0)
	// COLOR IT
	coinLabel := tui.NewLabel(fmt.Sprintf("Coin: %s", product))
	priceLabel := tui.NewLabel(fmt.Sprintf("Price: %s", tickPrice.Price))
	grid.AppendRow(tui.NewHBox(coinLabel, priceLabel))

	headerBar := tui.NewVBox(grid)
	headerBar.SetTitle(name)
	headerBar.SetBorder(true)

	// left and right side
	asksAdds := tui.NewVBox()
	asksAdds.SetSizePolicy(tui.Expanding, tui.Maximum)
	for _, m := range asks {
		asksAdds.Append(tui.NewHBox(
			tui.NewLabel(m.Price),
			tui.NewLabel(m.Volume),
		))
	}
	asksBox := tui.NewVBox(tui.NewVBox(
		tui.NewHBox(tui.NewLabel("Price"), tui.NewLabel("Volume"))),
		asksAdds)
	asksBox.SetTitle("Asks")
	asksBox.SetBorder(true)

	bidsAdds := tui.NewVBox()
	bidsAdds.SetSizePolicy(tui.Expanding, tui.Maximum)
	for _, m := range asks {
		bidsAdds.Append(tui.NewHBox(
			tui.NewLabel(m.Price),
			tui.NewLabel(m.Volume),
		))
	}
	bidsBox := tui.NewVBox(tui.NewVBox(
		tui.NewHBox(tui.NewLabel("Price"), tui.NewLabel("Volume"))),
		bidsAdds)
	bidsBox.SetTitle("Bids")
	bidsBox.SetBorder(true)

	// finish
	orderBookBox := tui.NewHBox(asksBox, bidsBox)
	orderBookBox.SetSizePolicy(tui.Preferred, tui.Expanding)
	root := tui.NewVBox(headerBar, orderBookBox)
	ui, err := tui.New(root)
	ui.SetKeybinding("Ctrl+C", func() { ui.Quit() })
	ui.SetKeybinding("ESC", func() { ui.Quit() })

	err = ui.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func getColor() {

}