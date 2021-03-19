package main

import (
	_ "fmt"
	_ "io/ioutil"
	_ "log"
	_ "os"
	_ "strconv"
	"strings"
	_ "time"

	"github.com/atsushi-kitazawa/gist_viewer/gist"
	"github.com/atsushi-kitazawa/gist_viewer/gui"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// const corporate = "aaa bbb the ccc ddd eee fff ggg"
const username = "atsushi-kitazawa"

var isTable bool = false

func main() {
	// get gist infomation.
	gistInfo := gist.NewGist(username)

	// initialization tview component.
	app := tview.NewApplication()
	gui.InitTview(app)
	gistList := gui.List
	description := gui.Description
	content := gui.Content
	flex := gui.Flex

	// list gist with table
	for i, g := range gistInfo {
		for _, f := range g.Files {
			gistList.SetCell(i, 0, &tview.TableCell{
				Text:            " " + f.Filename,
				NotSelectable:   false,
				Align:           tview.AlignLeft,
				Color:           tcell.ColorYellow,
				BackgroundColor: tcell.ColorDefault,
			})
		}
	}

	// table select func
	gistList.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			gistList.SetSelectable(true, false)
		}
	}).SetSelectedFunc(func(row int, col int) {
		row, col = gistList.GetSelection()
		cell := gistList.GetCell(row, col)
		description.SetText(gist.GetId(strings.Trim(cell.Text, " ")) + "\n" + gist.GetUrl(strings.Trim(cell.Text, " ")))
		content.SetText(gist.GetContent(gist.GetRawUrl(strings.Trim(cell.Text, " "))))
	})

	// set root
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

//numSelections := 0
//	go func() {
//		for _, word := range strings.Split(corporate, " ") {
//			if word == "the" {
//				word = "[#ff0000]the[white]"
//			}
//			if word == "to" {
//				word = fmt.Sprintf(`["%d"]to[""]`, numSelections)
//				numSelections++
//			}
//			//fmt.Fprintf(textView, "%s ", word)
//			time.Sleep(75 * time.Millisecond)
//		}
//	}()
