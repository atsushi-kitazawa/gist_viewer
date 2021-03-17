package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	_ "strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const corporate = "aaa bbb the ccc ddd eee fff ggg"

var isTable bool = false

func main() {
	app := tview.NewApplication()
	// create TextView
	textView := tview.NewTextView().
		SetDynamicColors(true).SetWordWrap(true).SetChangedFunc(func() { app.Draw() })
	textView.SetBorder(true).SetTitle("gist list")
	// create Table
	table := tview.NewTable().
		SetBorders(true).
		Select(0, 0).
		SetFixed(1, 1).
		SetSelectable(true, false)
	flex := tview.NewFlex().
		AddItem(textView, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("description"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("gist content"), 0, 10, false), 0, 2, false).
		AddItem(table, 0, 3, true)

	numSelections := 0
	go func() {
		for _, word := range strings.Split(corporate, " ") {
			if word == "the" {
				word = "[#ff0000]the[white]"
			}
			if word == "to" {
				word = fmt.Sprintf(`["%d"]to[""]`, numSelections)
				numSelections++
			}
			//fmt.Fprintf(textView, "%s ", word)
			time.Sleep(75 * time.Millisecond)
		}
	}()

	// set cell
	table.SetCell(0, 0, &tview.TableCell{
		Text:            " aaa",
		NotSelectable:   false,
		Align:           tview.AlignLeft,
		Color:           tcell.ColorYellow,
		BackgroundColor: tcell.ColorDefault,
	})
	table.SetCell(1, 0, &tview.TableCell{
		Text:            " aaa",
		NotSelectable:   false,
		Align:           tview.AlignLeft,
		Color:           tcell.ColorYellow,
		BackgroundColor: tcell.ColorDefault,
	})
	table.SetCell(2, 0, &tview.TableCell{
		Text:            " aaa",
		NotSelectable:   false,
		Align:           tview.AlignLeft,
		Color:           tcell.ColorYellow,
		BackgroundColor: tcell.ColorDefault,
	})

	// file list
	files := listFiles()
	for _, f := range files {
		fmt.Fprintf(textView, "%s \n", " "+f.Name())
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
