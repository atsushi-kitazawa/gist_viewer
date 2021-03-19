package main

import (
	_ "fmt"
	"github.com/atsushi-kitazawa/gist_viewer/gist"
	_ "io/ioutil"
	_ "log"
	_ "os"
	_ "strconv"
	"strings"
	_ "time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// const corporate = "aaa bbb the ccc ddd eee fff ggg"
const username = "atsushi-kitazawa"

var isTable bool = false

func main() {
	gistInfo := gist.NewGist(username)

	app := tview.NewApplication()
	// create textView
	textView := tview.NewTextView().
		SetDynamicColors(true).SetWordWrap(true).SetChangedFunc(func() { app.Draw() })
	textView.SetBorder(true).SetTitle("gist list")
	// create table
	table := tview.NewTable().
		Select(0, 0).
		SetFixed(1, 1).
		SetSelectable(true, false)
	table.SetBorder(true).SetTitle(" gist list")
	// create box
	description := tview.NewTextView().SetChangedFunc(func() { app.Draw() })
	description.SetBorder(true).SetTitle(" description")
	content := tview.NewTextView().SetChangedFunc(func() { app.Draw() })
	content.SetBorder(true).SetTitle(" content")
	// create flex
	flex := tview.NewFlex().
		AddItem(table, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(description, 0, 1, false).
			AddItem(content, 0, 10, false), 0, 2, false)

	// list gist with table
	for i, g := range gistInfo {
		for _, f := range g.Files {
			table.SetCell(i, 0, &tview.TableCell{
				Text:            " " + f.Filename,
				NotSelectable:   false,
				Align:           tview.AlignLeft,
				Color:           tcell.ColorYellow,
				BackgroundColor: tcell.ColorDefault,
			})
		}
	}

	// table select func
	table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			table.SetSelectable(true, false)
		}
	}).SetSelectedFunc(func(row int, col int) {
		row, col = table.GetSelection()
		cell := table.GetCell(row, col)
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
