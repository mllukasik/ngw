package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Repository struct {
	delegate git.Repository
	head     *plumbing.Reference
}

type Branch struct {
	Name    string
	RawName string
	Current bool
	Remote  bool
}

type PushOptions struct {
	Branch *string
	Remote *string
}
