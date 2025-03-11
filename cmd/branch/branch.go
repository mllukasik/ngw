package branch

import (
	"os"

	"github.com/mllukasik/ngw/git"
	"github.com/mllukasik/ngw/view/app"
	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "shows branches",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd.Help()
			return
		}
		app.NewApplication().BranchView().Run()
	},
}

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "removes all branches",
	Run: func(cmd *cobra.Command, args []string) {
		prune()
	},
}

func init() {
	BranchCmd.AddCommand(pruneCmd)
}

func prune() {
	repo, err := git.NewRepositoryWD()
	handleError(err, "prune")
	err = repo.PruneBranch()
	handleError(err, "prune")
}

func handleError(err error, process string) {
	if err == nil {
		return
	}
	os.Exit(1)
	println("Could not " + process + ": " + err.Error())
}
