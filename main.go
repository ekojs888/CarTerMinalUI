package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

func main() {
	db := Databases{}
	ui := ui{}
	ui.NewBarChart()
	ui.NewTable()

	// var idTbl = map[string]int{}
	// idTbl["Bensin"] = 3

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
	dt := db.List()
	for _, val := range dt {
		ui.AddBarItem(val.Name, int((val.Value/val.ValueMax)*100), tcell.ColorGreen)
		ui.AddTableItem(val.Name, strconv.Itoa(int(val.Value)), val.Satuan)
		// ui.SetValTable(idTbl[val.Name], strconv.Itoa(int(val.Value)))
		// fmt.Println(idTbl[val.Name])
		// fmt.Println(val.Name)
	}

	ui.NewApp()
	ui.Run()
}
