package main

import (
	"fmt"
	_ "fmt"
	"io/ioutil"
	"log"
	"os"
	_ "strconv"
	"strings"
	_ "time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// const corporate = "aaa bbb the ccc ddd eee fff ggg"

var isTable bool = false

func main() {
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
		//AddItem(textView, 0 , 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(description, 0, 1, false).
			AddItem(content, 0, 10, false), 0, 2, false)

	// list file with table
	files := listFiles()
	for i, f := range files {
		table.SetCell(i, 0, &tview.TableCell{
			Text:            " " + f.Name(),
			NotSelectable:   false,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorYellow,
			BackgroundColor: tcell.ColorDefault,
		})
	}

	// key bind
	// global
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			if isTable {
				app.SetFocus(table)
				isTable = !isTable
			} else {
				app.SetFocus(textView)
				isTable = !isTable
			}
		default:
			return event
		}
		return event
	})
	// textview
	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlN:
			textView.SetText(textView.GetText(true) + " ctrl+n")
		case tcell.KeyCtrlP:
			textView.SetText(textView.GetText(true) + " ctrl+p")
		case tcell.KeyEnter:
			textView.SetText(textView.GetText(true) + " return")
		default:
			return event
		}
		return event
	})

	// table select func
	table.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			table.SetSelectable(true, false)
		}
	}).SetSelectedFunc(func(row int, col int) {
		table.GetCell(row, col).SetTextColor(tcell.ColorRed)
		row, col = table.GetSelection()
		cell := table.GetCell(row, col)
		description.SetText(cell.Text)
		content.SetText(getContent(cell.Text))
	})

	// set root
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func listFiles() []os.FileInfo {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func getContent(file string) string {
	bytes, err := ioutil.ReadFile("./" + strings.Trim(file, " "))
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	return string(bytes)
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
