package cmd

import (
	"github.com/mllukasik/ngw/cmd/branch"
	"github.com/mllukasik/ngw/cmd/push"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ngw",
	Short: "Needless Git Wrapper",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(branch.BranchCmd)
	rootCmd.AddCommand(push.PushCmd)
}
