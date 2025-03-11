package git

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func NewRepositoryWD() (*Repository, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	return &Repository{
		delegate: *repo,
	}, nil
}

func (repo *Repository) PruneBranch() error {
	branches, err := repo.Branches()
	if err != nil {
		return err
	}
	for _, branch := range branches {
		if branch.Current {
			continue
		}
		err = repo.DeleteBranch(branch)
		if err != nil {
			break
		}
	}
	return err
}

func (repo *Repository) DeleteBranch(branch Branch) error {
	branchRef := referenceName(branch)
	err := repo.delegate.Storer.RemoveReference(branchRef)
	return err
}

func (repo *Repository) Push(options PushOptions) error {
	remote, err := repo.getDefaultRemote(options.Remote)
	if err != nil {
		return err
	}
	branch, err := repo.CurrentBranch()
	if err != nil {
		return err
	}

	//we cannot use git-go here because of credentials for now
	cmd := exec.Command("git", "push", remote, branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}

func (repo *Repository) CurrentBranch() (string, error) {
	if repo.head != nil {
		name := repo.head.Name().Short()
		return name, nil
	}
	head, err := repo.delegate.Head()
	if err != nil {
		return "", err
	}
	if head.Name().IsBranch() {
		name := head.Name().Short()
		repo.head = head
		return name, nil
	}
	return "", NgwGitError{message: "Could not determine current branch"}
}

func (repo *Repository) Checkout(branch Branch) error {
	//we cannot use git-go here because i dont know how to replacate checkout to non existing remote branch
	cmd := exec.Command("git", checkoutCommand(branch), branch.RawName)
	stderr := new(bytes.Buffer)
	cmd.Stderr = stderr
	err := cmd.Run()
	if stderr.Len() != 0 {
		return NgwGitError{message: stderr.String()}
	}
	return err
}

func checkoutCommand(branch Branch) string {
	if branch.Remote {
		return "switch"
	}
	return "checkout"
}

func (repo *Repository) Branches() ([]Branch, error) {
	refIter, err := repo.delegate.References()
	if err != nil {
		return []Branch{}, nil
	}
	references, err := storer.NewReferenceFilteredIter(
		func(r *plumbing.Reference) bool {
			return r.Name().IsBranch() || r.Name().IsRemote()
		}, refIter), nil

	if err != nil {
		return []Branch{}, nil
	}

	branches := []Branch{}
	references.ForEach(func(branch *plumbing.Reference) error {
		branches = append(branches, Branch{
			Name:    branch.Name().Short(),
			RawName: removePrefixes(branch.Name().Short()),
			Current: repo.isHead(branch),
			Remote:  branch.Name().IsRemote(),
		})
		return nil
	})
	return branches, nil
}

func removePrefixes(branchName string) string {
	index := strings.LastIndex(branchName, "/")
	if index == -1 {
		return branchName
	}

	return branchName[index+1:]
}

func (repo *Repository) isHead(branch *plumbing.Reference) bool {
	if repo.head == nil {
		repo.CurrentBranch()
	}
	return repo.head.Hash() == branch.Hash()
}

func (repo *Repository) getDefaultRemote(value *string) (string, error) {
	if value != nil {
		return *value, nil
	}
	remotes, err := repo.delegate.Remotes()
	if err != nil {
		return "", err
	}
	if len(remotes) != 1 {
		return "", NgwGitError{message: "There is no exactly one remote defined. Please specify remote"}
	}
	return remotes[0].Config().Name, nil
}

func referenceName(branch Branch) plumbing.ReferenceName {
	if branch.Remote {
		return plumbing.ReferenceName("refs/remotes/" + branch.Name)
	}
	return plumbing.NewBranchReferenceName(branch.Name)
}

type NgwGitError struct {
	message string
}

func (err NgwGitError) Error() string {
	return err.message
}
