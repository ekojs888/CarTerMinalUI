package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/navidys/tvxwidgets"
	"github.com/rivo/tview"
)

type ui struct {
	app *tview.Application
	// bdy *tview.Flex
	graph       *tvxwidgets.BarChart
	list        *tview.Table
	listMusic   *tview.List
	idList      int
	idListMusic int
	Music       *mocp
}

type mocp struct {
}

func (u *mocp) Start() {
	u.run("-S")
}

func (u *mocp) PlaylistAdd(path string) {
	u.run("-a", path)
}

func (u *mocp) PlaylistClear() {
	u.run("--clear")
}

func (u *mocp) PlayFirst() {
	u.run("--play")
}

func (u *mocp) GetCurentPlay() (filename string) {
	byt := u.run("-i")
	dt := strings.Split(string(byt), "\n")
	dt = strings.Split(dt[1], " ")
	return dt[1]
}

func (u *mocp) TogglePause() {
	u.run("--toggle-pause")
}

func (u *mocp) RunServer() {
	u.run("--server")
}

func (u *mocp) run(c ...string) (out []byte) {
	// fmt.Println(c)
	out, _ = exec.Command("/usr/bin/mocp", c...).Output()
	return out
}

func (u *ui) NewBarChart() {
	u.graph = tvxwidgets.NewBarChart()
	u.graph.SetRect(4, 2, 50, 20)
	u.graph.SetBorder(true)
	u.graph.SetTitle("System Resource Usage")
	u.graph.SetMaxValue(100)
}
func (u *ui) AddBarItem(name string, val int, color tcell.Color) {
	u.graph.AddBar(name, val, color)
}

func (u *ui) SetBarValue(name string, val int) {
	u.graph.SetBarValue(name, val)
}

func (u *ui) NewListMusic() {
	u.listMusic = tview.NewList()
	u.listMusic.ShowSecondaryText(false)
	u.listMusic.SetBorder(true)
	u.listMusic.SetTitle("Music")
	u.listMusic.SetBorderPadding(0, 0, 0, 0)

	// u.listMusic.AddItem("List item 1", "", '1', nil)
	// u.listMusic.AddItem("List item 2", "", '2', nil)
	// 		AddItem("List item 3", "Some explanatory text", 'c', nil).
	// 		AddItem("List item 4", "Some explanatory text", 'd', nil).
	// 		AddItem("Quit", "Press to exit", 'q', func() {
	// 			app.Stop()
	// 		})
}

func (u *ui) ReadDir(path string) []fs.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func (u *ui) AddItemMusicDir(value, dir string) {
	if value == ".." {
		u.listMusic.AddItem(value, dir, 0, func() {
			files := u.ReadDir(dir)
			u.listMusic.Clear()
			u.idListMusic = 0
			for _, file := range files {
				if file.Name()[0:1] == "." {
					continue
				}
				if file.IsDir() {
					u.AddItemMusicDir(file.Name(), dir)
				} else {
					u.AddItemMusic(file.Name(), dir)
				}
			}

		})
	} else {
		u.idListMusic += 1
		u.listMusic.AddItem(strconv.Itoa(u.idListMusic)+". "+value, dir+"/"+value, 0, func() {
			// path := dir + "/" + value
			ci := u.listMusic.GetCurrentItem()
			_, path := u.listMusic.GetItemText(ci)

			u.Music.PlaylistClear()
			u.Music.PlaylistAdd(path)
			// u.Music.PlayFirst()

			files := u.ReadDir(path)
			u.listMusic.Clear()
			u.idListMusic = 0
			u.AddItemMusicDir("..", dir)
			for _, file := range files {
				if file.Name()[0:1] == "." {
					continue
				}
				if file.IsDir() {
					u.AddItemMusicDir(file.Name(), path)
				} else {
					u.AddItemMusic(file.Name(), path)
				}
			}

		})
	}
	// path := dir + "/" + value
	// file, _ := os.Open(path)
	// defer file.Close()
	// fileInfo, _ := file.Stat()
	// if fileInfo.IsDir() {
	// 	u.Music.PlaylistClear()
	// 	u.Music.PlaylistAdd(path)
	// }
}
func (u *ui) AddItemMusic(value string, path string) {
	u.idListMusic += 1
	u.listMusic.AddItem(strconv.Itoa(u.idListMusic)+". "+value, path+"/"+value, 0, func() {
		ci := u.listMusic.GetCurrentItem()
		_, txt := u.listMusic.GetItemText(ci)
		u.Music.run("-l", txt)
	})
}

func (u *ui) NewTable() {
	u.list = tview.NewTable()
	u.list.SetBorder(true)
	u.list.SetBorders(false)
	u.list.SetTitle("List")
	u.list.SetBorderPadding(1, 0, 0, 0)

	// u.list.SetCell(0, 0, tview.NewTableCell("1"))
	// u.list.SetCell(0, 1, tview.NewTableCell("AC"))
	// u.list.SetCell(0, 3, tview.NewTableCell(":"))
	// u.list.SetCell(0, 4, tview.NewTableCell("OFF"))

	// u.list.SetCell(1, 0, tview.NewTableCell("2"))
	// u.list.SetCell(1, 1, tview.NewTableCell("Suhu"))
	// u.list.SetCell(1, 3, tview.NewTableCell(":"))
	// u.list.SetCell(1, 4, tview.NewTableCell("26 °C"))

	// u.list.SetCell(2, 0, tview.NewTableCell("3"))
	// u.list.SetCell(2, 1, tview.NewTableCell("Mesin"))
	// u.list.SetCell(2, 3, tview.NewTableCell(":"))
	// u.list.SetCell(2, 4, tview.NewTableCell("7000 Rpm"))

	// u.list.SetCell(2, 0, tview.NewTableCell("3"))
	// u.list.SetCell(2, 1, tview.NewTableCell("Mesin"))
	// u.list.SetCell(2, 3, tview.NewTableCell(":"))
	// u.list.SetCell(2, 4, tview.NewTableCell("7000 Rpm"))

	// u.list.SetCell(3, 0, tview.NewTableCell("4"))
	// u.list.SetCell(3, 1, tview.NewTableCell("Bensin"))
	// u.list.SetCell(3, 3, tview.NewTableCell(":"))
	// u.list.SetCell(3, 4, tview.NewTableCell("25 Liter"))
}

func (u *ui) AddTableItem(name, val, satuan string) {
	u.list.SetCell(u.idList, 0, tview.NewTableCell(strconv.Itoa(u.idList+1)))
	u.list.SetCell(u.idList, 1, tview.NewTableCell(name))
	u.list.SetCell(u.idList, 3, tview.NewTableCell(":"))
	u.list.SetCell(u.idList, 4, tview.NewTableCell(val))
	u.list.SetCell(u.idList, 5, tview.NewTableCell(satuan))
	u.idList++
}

func (u *ui) SetValTable(row int, val string) {
	// update list table
	u.list.GetCell(row, 4).SetText(val)
}

func (u *ui) NewApp() {
	// u.graph.SetRect(4, 2, 50, 20)
	// u.graph.SetBorder(true)
	// u.graph.SetTitle("System Resource Usage")
	// display system metric usage
	// u.graph.AddBar("mesin", 50, tcell.ColorBlue)
	// u.graph.AddBar("suhu", 25, tcell.ColorRed)
	// u.graph.AddBar("bensin", 80, tcell.ColorGreen)
	// u.graph.AddBar("aki", 12, tcell.ColorOrange)
	// u.graph.SetMaxValue(100)

	// update list table
	// list.GetCell(1, 4).SetText("40 °C")

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
			AddItem(u.graph, 0, 1, false),
		// AddItem(gauge1, 4, 1, false),
		0, 1, false)
	flex.AddItem(
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			// AddItem(gauge, 3, 1, false).
			// AddItem(gauge1, 3, 1, false).
			AddItem(u.list, 0, 1, false).
			AddItem(u.listMusic, 0, 1, true),
		0, 1, false)
	flex.GetItem(0).Blur()
	u.app.SetRoot(flex, true)
	u.app.SetFocus(u.listMusic)
}

func (u *ui) Run() {
	// u.app.EnableMouse(true)
	if err := u.app.Run(); err != nil {
		panic(err)
	}
}

// func uis() {
// 	app := tview.NewApplication()

// 	barGraph := tvxwidgets.NewBarChart()
// 	barGraph.SetRect(4, 2, 50, 20)
// 	barGraph.SetBorder(true)
// 	barGraph.SetTitle("System Resource Usage")
// 	// display system metric usage
// 	barGraph.AddBar("cpu", 80, tcell.ColorBlue)
// 	barGraph.AddBar("mem", 20, tcell.ColorRed)
// 	barGraph.AddBar("swap", 40, tcell.ColorGreen)
// 	barGraph.AddBar("disk", 40, tcell.ColorOrange)
// 	barGraph.SetMaxValue(100)

// 	gauge := tvxwidgets.NewPercentageModeGauge()
// 	gauge.SetTitle("percentage mode gauge")
// 	gauge.SetRect(10, 4, 50, 3)
// 	gauge.SetBorder(true)
// 	gauge.SetMaxValue(50)

// 	gauge1 := tvxwidgets.NewUtilModeGauge()
// 	gauge1.SetLabel("cpu usage:")
// 	gauge1.SetLabelColor(tcell.ColorLightSkyBlue)
// 	gauge1.SetRect(10, 4, 50, 3)
// 	gauge1.SetWarnPercentage(65)
// 	gauge1.SetCritPercentage(80)
// 	gauge1.SetBorder(true)

// 	box := tview.NewBox()
// 	box.SetBorder(true)
// 	box.SetTitle("Hello, world!")
// 	box.SetBackgroundColor(0)
// 	box.SetBorderColor(tcell.Color109)

// 	list := tview.NewList().
// 		AddItem("List item 1", "Some explanatory text", 'a', nil).
// 		AddItem("List item 2", "Some explanatory text", 'b', nil).
// 		AddItem("List item 3", "Some explanatory text", 'c', nil).
// 		AddItem("List item 4", "Some explanatory text", 'd', nil).
// 		AddItem("Quit", "Press to exit", 'q', func() {
// 			app.Stop()
// 		})
// 	list.SetBorder(true)

// 	flex := tview.NewFlex().
// 		AddItem(box, 0, 1, false).
// 		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
// 			AddItem(barGraph, 0, 1, false).
// 			AddItem(gauge, 4, 5, false).
// 			AddItem(gauge1, 4, 1, false),
// 			0, 2, false).
// 		AddItem(list, 0, 1, false)

// 	flex.GetItem(0).Blur()

// 	app.SetRoot(flex, true)
// 	app.EnableMouse(true)
// 	app.SetFocus(list)

// 	var count int = 0

// 	update := func() {
// 		// tick := time.NewTicker(500 * time.Millisecond)
// 		for {
// 			time.Sleep(1 * time.Second)

// 			if count > gauge.GetMaxValue() {
// 				count = 0
// 			} else {
// 				count = count + 1
// 			}
// 			gauge.SetValue(count)
// 			barGraph.SetBarValue("cpu", count)
// 			gauge1.SetValue(float64(count))
// 			app.Draw()
// 		}
// 	}
// 	go update()

// 	// go func() {
// 	// 	for {
// 	// 		time.Sleep(1 * time.Second)
// 	// 		countStr := strconv.Itoa(count)
// 	// 		app.QueueUpdateDraw(func() {
// 	// 			box.SetTitle("hallow " + countStr)
// 	// 			barGraph.SetBarValue("cpu", count)
// 	// 		})
// 	// 		count++
// 	// 	}

// 	// }()

// 	if err := app.Run(); err != nil {
// 		panic(err)
// 	}
// }
