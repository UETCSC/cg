package cmd

import (
	"cg/pkg/cmdutil"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CheckCmd)
}

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Verify whether the current question directory and content meet the standards.",
	Long:  "Verify whether the current question directory and content meet the standards.",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.Check()
	},
}
