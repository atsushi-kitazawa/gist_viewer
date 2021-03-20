package keybinds

import (
	"github.com/atsushi-kitazawa/gist_viewer/gui"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func SetGlobal(app *tview.Application) {
    // key bind
	// global
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
		    gui.MoveFocus(app)
		default:
			return event
		}
		return event
	})
	// textview
	//textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	//	switch event.Key() {
	//	case tcell.KeyCtrlN:
	//		textView.SetText(textView.GetText(true) + " ctrl+n")
	//	case tcell.KeyCtrlP:
	//		textView.SetText(textView.GetText(true) + " ctrl+p")
	//	case tcell.KeyEnter:
	//		textView.SetText(textView.GetText(true) + " return")
	//	default:
	//		return event
	//	}
	//	return event
	//})

}
