package main

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

func main() {
	db := Databases{}
	ui := ui{}
	ui.NewBarChart()
	ui.NewTable()

	var idTbl = map[string]int{}
	idTbl["Bensin"] = 3

	db.ConSqlLite()
	// db.Init()
	dt := db.List()
	for _, val := range dt {
		ui.AddBarItem(val.Name, int(val.Value), tcell.ColorGreen)
		ui.SetValTable(idTbl[val.Name], strconv.Itoa(int(val.Value)))
		fmt.Println(idTbl[val.Name])
		// fmt.Println(val.Name)
	}

	ui.NewApp()
	ui.Run()
}
