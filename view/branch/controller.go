package branch

import (
	"github.com/mllukasik/ngw/git"
)

type controller struct {
	repo     git.Repository
	branches []git.Branch
}

func newController() (*controller, error) {
	repo, err := git.NewRepositoryWD()
	if err != nil {
		return nil, err
	}
	controller := &controller{
		repo: *repo,
	}
	err = controller.refresh()
	if err != nil {
		return nil, err
	}
	return controller, err
}

func (controller *controller) refresh() error {
	branches, err := controller.repo.Branches()
	if err != nil {
		return err
	}
	controller.branches = branches
	return nil
}

func (controller controller) checkout(index int) (*git.Branch, error) {
	branch := controller.currentBranch(index)
	if branch == nil {
		return nil, nil
	}
	err := controller.repo.Checkout(*branch)
	return branch, err
}

func (controller controller) deleteBranch(index int) (*git.Branch, error) {
	branch := controller.currentBranch(index)
	if branch == nil {
		return nil, nil
	}
	err := controller.repo.DeleteBranch(*branch)
	if err != nil {
		return branch, err
	}
	controller.refresh()
	return branch, err
}

func (controller controller) currentBranch(index int) *git.Branch {
	length := len(controller.branches)
	if index < 0 || index >= length {
		return nil
	}
	return &controller.branches[index]
}
