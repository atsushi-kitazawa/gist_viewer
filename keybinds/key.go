package keybinds

import (
	"github.com/rivo/tview"
	_ "github.com/gdamore/tcell"
)

var isTable = false

func SetGlobal(app tview.Application) {
    // key bind
	// global
	//app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
	//	switch event.Key() {
	//	case tcell.KeyTab:
	//		if isTable {
	//			app.SetFocus(table)
	//			isTable = !isTable
	//		} else {
	//			app.SetFocus(textView)
	//			isTable = !isTable
	//		}
	//	default:
	//		return event
	//	}
	//	return event
	//})
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
