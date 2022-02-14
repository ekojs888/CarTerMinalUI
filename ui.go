package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

type ui struct {
	app *tview.Application
	// bdy *tview.Flex
}

func (u *ui) NewApp() {
	barGraph := tvxwidgets.NewBarChart()
	barGraph.SetRect(4, 2, 50, 20)
	barGraph.SetBorder(true)
	barGraph.SetTitle("System Resource Usage")
	// display system metric usage
	barGraph.AddBar("mesin", 50, tcell.ColorBlue)
	barGraph.AddBar("suhu", 25, tcell.ColorRed)
	barGraph.AddBar("bensin", 80, tcell.ColorGreen)
	barGraph.AddBar("aki", 12, tcell.ColorOrange)
	barGraph.SetMaxValue(100)

	list := tview.NewTable()
	list.SetBorder(true)
	list.SetBorders(false)
	list.SetTitle("List")
	list.SetBorderPadding(1, 0, 0, 0)

	list.SetCell(0, 0, tview.NewTableCell("1"))
	list.SetCell(0, 1, tview.NewTableCell("AC"))
	list.SetCell(0, 3, tview.NewTableCell(":"))
	list.SetCell(0, 4, tview.NewTableCell("OFF"))

	list.SetCell(1, 0, tview.NewTableCell("2"))
	list.SetCell(1, 1, tview.NewTableCell("Suhu"))
	list.SetCell(1, 3, tview.NewTableCell(":"))
	list.SetCell(1, 4, tview.NewTableCell("26 °C"))

	list.SetCell(2, 0, tview.NewTableCell("3"))
	list.SetCell(2, 1, tview.NewTableCell("Mesin"))
	list.SetCell(2, 3, tview.NewTableCell(":"))
	list.SetCell(2, 4, tview.NewTableCell("7000 Rpm"))

	list.SetCell(2, 0, tview.NewTableCell("3"))
	list.SetCell(2, 1, tview.NewTableCell("Mesin"))
	list.SetCell(2, 3, tview.NewTableCell(":"))
	list.SetCell(2, 4, tview.NewTableCell("7000 Rpm"))

	list.SetCell(3, 0, tview.NewTableCell("4"))
	list.SetCell(3, 1, tview.NewTableCell("Bensin"))
	list.SetCell(3, 3, tview.NewTableCell(":"))
	list.SetCell(3, 4, tview.NewTableCell("25 Liter"))

	// update list table
	list.GetCell(1, 4).SetText("40 °C")

	gauge := tvxwidgets.NewPercentageModeGauge()
	gauge.SetTitle("percentage mode gauge")
	gauge.SetRect(10, 4, 50, 3)
	gauge.SetBorder(true)
	gauge.SetMaxValue(100)
	gauge.SetValue(75)

	gauge1 := tvxwidgets.NewUtilModeGauge()
	gauge1.SetLabel("cpu usage:")
	gauge1.SetLabelColor(tcell.ColorLightSkyBlue)
	gauge1.SetRect(10, 4, 50, 3)
	gauge1.SetWarnPercentage(65)
	gauge1.SetCritPercentage(80)
	gauge1.SetBorder(true)

	u.app = tview.NewApplication()
	flex := tview.NewFlex()
	flex.AddItem(
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(barGraph, 0, 1, false),
		// AddItem(gauge1, 4, 1, false),
		0, 2, false)
	flex.AddItem(
		tview.NewFlex().
			// SetDirection(tview.FlexRow).
			// AddItem(gauge, 3, 1, false).
			// AddItem(gauge1, 3, 1, false).
			AddItem(list, 0, 1, true),
		0, 1, false)
	flex.GetItem(0).Blur()
	u.app.SetRoot(flex, true)
	// u.app.SetFocus(list)
}

func (u *ui) Run() {
	// u.app.EnableMouse(true)
	if err := u.app.Run(); err != nil {
		panic(err)
	}
}

func uis() {
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
