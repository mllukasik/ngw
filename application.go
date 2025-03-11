package main

import (
	"github.com/rivo/tview"
)

func RunApplication() Application {
	return Application{
		tview.NewApplication(),
	}
}

type AppView struct {
	View tview.Primitive
}

type Application struct {
	delegate *tview.Application
}
