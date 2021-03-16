package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	_ "strconv"
	"strings"
	"time"

	_ "github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const corporate = "aaa bbb the ccc ddd eee fff ggg"

func main() {
	app := tview.NewApplication()
	// create TextView
	textView := tview.NewTextView().
	    SetDynamicColors(true).SetWordWrap(true).SetChangedFunc(func() { app.Draw() })
	textView.SetBorder(true).SetTitle("gist list")
	flex := tview.NewFlex().
		AddItem(textView, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("description"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("gist content"), 0, 10, false), 0, 2, false)

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

	// file list
	files := listFiles()
	for _, f := range files {
	    fmt.Fprintf(textView, "%s \n", " " + f.Name())
	}

	// key bind

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
