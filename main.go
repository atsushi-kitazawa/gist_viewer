package main

import (
	"flag"
	_ "fmt"
	_ "io/ioutil"
	_ "log"
	_ "os"
	_ "strconv"
	"strings"
	_ "time"

	"github.com/atsushi-kitazawa/gist_viewer/gist"
	"github.com/atsushi-kitazawa/gist_viewer/gui"
	"github.com/atsushi-kitazawa/gist_viewer/keybinds"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var isTable bool = false

func main() {
	// argument
	username := flag.String("u", "", "gist view target username")
	flag.Parse()

	// get gist infomation.
	gistInfo := gist.NewGist(*username)

	// initialization tview component.
	app := tview.NewApplication()
	gui.InitTview(app)
	gistList := gui.List
	description := gui.Description
	content := gui.Content
	flex := gui.Flex

	// set key binds
	keybinds.SetGlobal(app)

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
		description.SetText(gist.GetId(trimSpace(cell.Text)) + "\n" + gist.GetUrl(trimSpace(cell.Text)))
		content.SetText(gist.GetContent(gist.GetRawUrl(trimSpace(cell.Text))))
	})

	// set root
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func trimSpace(s string) string {
    return strings.Trim(s, " ")
}
