package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	barGraph := tvxwidgets.NewBarChart()
	barGraph.SetRect(4, 2, 50, 20)
	barGraph.SetBorder(true)
	barGraph.SetTitle("System Resource Usage")
	// display system metric usage
	barGraph.AddBar("cpu", 80, tcell.ColorBlue)
	barGraph.AddBar("mem", 20, tcell.ColorRed)
	barGraph.AddBar("swap", 40, tcell.ColorGreen)
	barGraph.AddBar("disk", 40, tcell.ColorOrange)
	barGraph.SetMaxValue(100)

	gauge := tvxwidgets.NewPercentageModeGauge()
	gauge.SetTitle("percentage mode gauge")
	gauge.SetRect(10, 4, 50, 3)
	gauge.SetBorder(true)
	gauge.SetMaxValue(50)

	gauge1 := tvxwidgets.NewUtilModeGauge()
	gauge1.SetLabel("cpu usage:")
	gauge1.SetLabelColor(tcell.ColorLightSkyBlue)
	gauge1.SetRect(10, 4, 50, 3)
	gauge1.SetWarnPercentage(65)
	gauge1.SetCritPercentage(80)
	gauge1.SetBorder(true)

	box := tview.NewBox()
	box.SetBorder(true)
	box.SetTitle("Hello, world!")
	box.SetBackgroundColor(0)
	box.SetBorderColor(tcell.Color109)

	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	list.SetBorder(true)

	flex := tview.NewFlex().
		AddItem(box, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(barGraph, 0, 1, false).
			AddItem(gauge, 4, 5, false).
			AddItem(gauge1, 4, 1, false),
			0, 2, false).
		AddItem(list, 0, 1, false)

	flex.GetItem(0).Blur()

	app.SetRoot(flex, true)
	app.EnableMouse(true)
	app.SetFocus(list)

	var count int = 0

	update := func() {
		// tick := time.NewTicker(500 * time.Millisecond)
		for {
			time.Sleep(1 * time.Second)

			if count > gauge.GetMaxValue() {
				count = 0
			} else {
				count = count + 1
			}
			gauge.SetValue(count)
			barGraph.SetBarValue("cpu", count)
			gauge1.SetValue(float64(count))
			app.Draw()
		}
	}
	go update()

	// go func() {
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		countStr := strconv.Itoa(count)
	// 		app.QueueUpdateDraw(func() {
	// 			box.SetTitle("hallow " + countStr)
	// 			barGraph.SetBarValue("cpu", count)
	// 		})
	// 		count++
	// 	}

	// }()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
