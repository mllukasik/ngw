package branch

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/mllukasik/ngw/view"
	"github.com/rivo/tview"
)

type branchView struct {
	controller   controller
	exitCallback func()
	scaffolld    *tview.Flex
	currentIndex func() int
}

func (branchView branchView) View() tview.Primitive {
	return branchView.scaffolld
}

func NewBranchView(exitCallback func()) view.AppView {
	controller, err := newController()
	if err != nil {
		fmt.Println("Could not init branch view: " + err.Error())
		exitCallback()
	}
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	view := branchView{
		controller:   *controller,
		exitCallback: exitCallback,
		scaffolld:    flex,
		currentIndex: func() int { return -1 },
	}
	view.build()
	return view
}

func (branchView branchView) build() {
	scaffolld := branchView.scaffolld
	scaffolld.Clear().
		AddItem(branchView.menuView(), 7, 1, false).
		AddItem(branchView.branchList(), 0, 1, true).
		SetInputCapture(branchView.inputCapture)
}

func (branchView *branchView) branchList() tview.Primitive {
	list := tview.NewList()
	branches := branchView.controller.branches
	for index, element := range branches {
		rune := rune(strconv.Itoa(index)[0])
		secondaryText := ""
		if element.Current {
			secondaryText = "current"
		}
		list.AddItem(element.Name, secondaryText, rune, nil)
	}
	list.SetBorder(true).SetTitle(getTitle(len(branches))).SetInputCapture(vimMotionForList(list))
	branchView.currentIndex = list.GetCurrentItem
	return list
}

func (branchView branchView) menuView() tview.Primitive {
	return tview.NewGrid().
		//col - 1
		AddItem(label(""), 0, 0, 1, 1, 0, 0, false).
		AddItem(label(""), 1, 0, 1, 1, 0, 0, false).
		AddItem(label(""), 2, 0, 1, 1, 0, 0, false).
		AddItem(label(""), 3, 0, 1, 1, 0, 0, false).
		AddItem(label(""), 4, 0, 1, 1, 0, 0, false).
		AddItem(label(""), 5, 0, 1, 1, 0, 0, false).
		AddItem(label(""), 6, 0, 1, 1, 0, 0, false).
		//col - 2
		AddItem(dlabel("<q>", "exit"), 0, 1, 1, 1, 0, 0, false).
		AddItem(label(""), 1, 1, 1, 1, 0, 0, false).
		AddItem(label(""), 2, 1, 1, 1, 0, 0, false).
		AddItem(label(""), 3, 1, 1, 1, 0, 0, false).
		AddItem(label(""), 4, 1, 1, 1, 0, 0, false).
		AddItem(label(""), 5, 1, 1, 1, 0, 0, false).
		AddItem(label(""), 6, 1, 1, 1, 0, 0, false).
		//col - 3
		AddItem(dlabel("<c>", "checkout"), 0, 2, 1, 1, 0, 0, false).
		AddItem(dlabel("<D>", "delete"), 1, 2, 1, 1, 0, 0, false).
		AddItem(label(""), 2, 2, 1, 1, 0, 0, false).
		AddItem(label(""), 3, 2, 1, 1, 0, 0, false).
		AddItem(label(""), 4, 2, 1, 1, 0, 0, false).
		AddItem(label(""), 5, 2, 1, 1, 0, 0, false).
		AddItem(label(""), 6, 2, 1, 1, 0, 0, false).
		//col - 4
		AddItem(label(""), 0, 3, 1, 1, 0, 0, false).
		AddItem(label(""), 1, 3, 1, 1, 0, 0, false).
		AddItem(label(""), 2, 3, 1, 1, 0, 0, false).
		AddItem(label(""), 3, 3, 1, 1, 0, 0, false).
		AddItem(label(""), 4, 3, 1, 1, 0, 0, false).
		AddItem(label(""), 5, 3, 1, 1, 0, 0, false).
		AddItem(label(""), 6, 3, 1, 1, 0, 0, false).
		//col - 5
		AddItem(label(""), 0, 4, 1, 1, 0, 0, false).
		AddItem(label(""), 1, 4, 1, 1, 0, 0, false).
		AddItem(label(""), 2, 4, 1, 1, 0, 0, false).
		AddItem(label(""), 3, 4, 1, 1, 0, 0, false).
		AddItem(label(""), 4, 4, 1, 1, 0, 0, false).
		AddItem(label(""), 5, 4, 1, 1, 0, 0, false).
		AddItem(label(""), 6, 4, 1, 1, 0, 0, false).
		//col - 6
		AddItem(label(""), 0, 5, 1, 1, 0, 0, false).
		AddItem(label(""), 1, 5, 1, 1, 0, 0, false).
		AddItem(label(""), 2, 5, 1, 1, 0, 0, false).
		AddItem(label(""), 3, 5, 1, 1, 0, 0, false).
		AddItem(label(""), 4, 5, 1, 1, 0, 0, false).
		AddItem(label(""), 5, 5, 1, 1, 0, 0, false).
		AddItem(label(""), 6, 5, 1, 1, 0, 0, false)
}

func dlabel(key string, text string) tview.Primitive {
	return tview.NewGrid().
		AddItem(clabel(key, tcell.ColorDeepPink), 0, 0, 1, 1, 0, 0, false).
		AddItem(clabel(text, tcell.ColorDefault), 0, 1, 1, 1, 0, 0, false)
}

func clabel(text string, color tcell.Color) tview.Primitive {
	return tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(text).SetTextColor(color)
}

func label(text string) tview.Primitive {
	return tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(text)
}

func getTitle(count int) string {
	title := fmt.Sprintf("Branches(all)[%d]", count)
	return title
}

func (branchView branchView) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'q':
		branchView.exitCallback()
		return nil
	case 'c':
		branchView.checkout()
		return nil
	case 'D':
		branchView.deleteBranch()
		return nil
	}
	return event
}

func (branchView branchView) checkout() {
	branch, err := branchView.controller.checkout(branchView.currentIndex())
	branchView.exitCallback()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Switched to branch '%s'\n", branch.RawName)
}
func (branchView branchView) deleteBranch() {
	branch, err := branchView.controller.deleteBranch(branchView.currentIndex())
	branchView.exitCallback()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Deleted branch '%s'\n", branch.RawName)
}

func vimMotionForList(list *tview.List) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'j':
			list.SetCurrentItem(list.GetCurrentItem() + 1)
			return nil
		case 'k':
			list.SetCurrentItem(list.GetCurrentItem() - 1)
			return nil
		}
		return event
	}
}
