package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/mllukasik/ngw/build"
	"github.com/mllukasik/ngw/cmd/branch"
	"github.com/mllukasik/ngw/cmd/push"
	"github.com/spf13/cobra"
)

func longDesc() string {
	var builder strings.Builder
	builder.WriteString("Needless Git Wrapper\n")
	builder.WriteString(fmt.Sprintf("Version: %s\n", build.Build().Version))
	if build.Build().Debug {
		builder.WriteString(fmt.Sprintf("Build date: %s\n", build.Build().Date.Format(time.ANSIC)))
	}
	return builder.String()
}

var rootCmd = &cobra.Command{
	Use:   "ngw",
	Short: "Needless Git Wrapper",
	Long:  longDesc(),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	Version: build.Build().Version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(branch.BranchCmd)
	rootCmd.AddCommand(push.PushCmd)
}
