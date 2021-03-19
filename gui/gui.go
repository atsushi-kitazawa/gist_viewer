package gui

import "github.com/rivo/tview"

var (
	Flex        *GistFlex        = &GistFlex{}
	List        *GistList        = &GistList{}
	Description *GistDescription = &GistDescription{}
	Content     *GistContent     = &GistContent{}
)

type GistFlex struct {
	*tview.Flex
}

type GistList struct {
	*tview.Table
}

type GistDescription struct {
	*tview.TextView
}

type GistContent struct {
	*tview.TextView
}

func InitTview(app *tview.Application) {
	List = &GistList{tview.NewTable().
		Select(0, 0).
		SetFixed(1, 1).
		SetSelectable(true, false)}
	List.SetBorder(true).SetTitle(" gist list")

	Description = &GistDescription{tview.NewTextView().SetChangedFunc(func() { app.Draw() })}
	Description.SetBorder(true).SetTitle(" description")

	Content = &GistContent{tview.NewTextView().SetChangedFunc(func() { app.Draw() })}
	Content.SetBorder(true).SetTitle(" content")

	Flex = &GistFlex{tview.NewFlex().
		AddItem(List, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(Description, 0, 1, false).
			AddItem(Content, 0, 10, false), 0, 2, false)}
	_ = Flex
}
