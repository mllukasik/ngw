package view

import (
	"github.com/rivo/tview"
)

type AppView interface {
	View() tview.Primitive
}
