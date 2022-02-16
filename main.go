package main

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	db := Databases{}
	ui := ui{}
	ui.NewBarChart()
	ui.NewTable()

	db.ConSqlLite()
	// db.Init()
	// db.Inserts(
	// 	&[]TblData{
	// 		{
	// 			Name:     "Bensin",
	// 			Value:    20,
	// 			ValueMax: 40,
	// 			Satuan:   "Liter",
	// 		},
	// 		{
	// 			Name:     "Mesin",
	// 			Value:    7000,
	// 			ValueMax: 10000,
	// 			Satuan:   "Rpm",
	// 		},
	// 		{
	// 			Name:     "Aki",
	// 			Value:    12,
	// 			ValueMax: 13,
	// 			Satuan:   "Volt",
	// 		},
	// 	},
	// )
	// db.Init()
	db.UpdateByName("Bensin", 15)

	go func() {
		var a float64 = 0
		for {
			time.Sleep(100 * time.Millisecond)
			db.UpdateByName("Bensin", a)
			if a >= 40 {
				a = 0
			} else {
				a++
			}

			dt := db.List()
			for _, val := range dt {
				ui.SetBarValue(val.Name, int((val.Value/val.ValueMax)*100))
				ui.SetValTable(int(val.ID)-1, strconv.Itoa(int(val.Value)))
			}
			ui.app.Draw()
		}
	}()

	dt := db.List()
	for _, val := range dt {
		ui.AddBarItem(val.Name, int((val.Value/val.ValueMax)*100), tcell.ColorGreen)
		ui.AddTableItem(val.Name, strconv.Itoa(int(val.Value)), val.Satuan)
	}

	ui.NewApp()
	ui.Run()
}
