package app

import (
	"github.com/mllukasik/ngw/view/branch"
	"github.com/rivo/tview"
)

func NewApplication() application {
	return application{
		tview.NewApplication(),
	}
}

type application struct {
	delegate *tview.Application
}

func (app application) BranchView() application {
	view := branch.NewBranchView(app.exit)
	scaffold := view.View()
	app.delegate.SetRoot(scaffold, true)
	app.delegate.SetFocus(scaffold)
	return app
}

func (app application) Run() application {
	if err := app.delegate.Run(); err != nil {
		panic(err)
	}
	return app
}

func (app application) exit() {
	app.delegate.Stop()
}
