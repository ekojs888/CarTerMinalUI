package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

var curDir string = ""

func readFlasdisk(ui ui) {

	dirf := "/home/eko/Downloads/fdiskcek/dirfolder.txt"
	dirr, err := os.Open(dirf)
	if err != nil {
		log.Fatal(err)
	}
	filescanner := bufio.NewScanner(dirr)
	filescanner.Split(bufio.ScanLines)
	var fileLines []string
	for filescanner.Scan() {
		fileLines = append(fileLines, filescanner.Text())
	}
	dirr.Close()

	dir := fileLines[0]

	if dir == "" {
		ui.listMusic.Clear()
		curDir = ""
	}

	if dir != "" && curDir != dir {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		ui.Music.Start()

		for _, file := range files {
			if file.IsDir() {
				if file.Name()[0:1] != "." {
					ui.AddItemMusicDir(file.Name(), dir)
				}
			} else {
				ui.AddItemMusic(file.Name(), dir)
			}
		}
		curDir = dir
	}
}

func main() {
	db := Databases{}
	ui := ui{}
	ui.NewBarChart()
	ui.NewTable()
	ui.NewListMusic()

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
			time.Sleep(30 * time.Millisecond)
			db.UpdateByName("Bensin", a)
			db.UpdateByName("Mesin", a/40*10000)
			db.UpdateByName("Aki", a/40*13)
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

			readFlasdisk(ui)

			ui.app.Draw()

		}
	}()

	color := []tcell.Color{
		tcell.ColorBlue,
		tcell.ColorRed,
		tcell.ColorGreen,
		tcell.ColorYellow,
		tcell.ColorWhite,
		tcell.ColorOrange,
	}
	dt := db.List()
	for key, val := range dt {
		ui.AddBarItem(val.Name, int((val.Value/val.ValueMax)*100), color[key])
		ui.AddTableItem(val.Name, strconv.Itoa(int(val.Value)), val.Satuan)
	}

	ui.NewApp()
	ui.Run()
}
