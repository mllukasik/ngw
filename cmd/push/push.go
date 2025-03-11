package push

import (
	"fmt"
	"os"

	"github.com/mllukasik/ngw/git"
	"github.com/spf13/cobra"
)

var PushCmd = &cobra.Command{
	Use:   "push",
	Short: "push current branch",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("specifing branch or remote is not supported yet")
			return
		}
		push()
	},
}

func push() {
	repo, err := git.NewRepositoryWD()
	handleError(err, "push")
	err = repo.Push(git.PushOptions{})
	handleError(err, "push")
}

func handleError(err error, process string) {
	if err == nil {
		return
	}
	os.Exit(1)
	println("Could not " + process + ": " + err.Error())
}
