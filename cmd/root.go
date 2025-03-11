package cmd

import (
	"github.com/mllukasik/ngw/cmd/branch"
	"github.com/mllukasik/ngw/cmd/push"
	"github.com/spf13/cobra"
)

func rootCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "ngw",
		Short: "Needless Git Wrapper",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		Version: version,
	}

}

func Execute(version string) error {
	var root = rootCmd(version)
	root.AddCommand(branch.BranchCmd)
	root.AddCommand(push.PushCmd)
	return root.Execute()
}
