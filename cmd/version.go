package cmd

import (
	"cg/pkg/util"
	"fmt"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Check cg version and update time.",
	Long:  "Check cg version and update time.",
	Run: func(cmd *cobra.Command, args []string) {
		Cyan := color.FgCyan.Render
		fmt.Println("Version:   ", Cyan(util.Version))
		fmt.Println("Build Time: ", Cyan(util.BuildTime))
		fmt.Println("Commit: ", Cyan(util.GitCommitId))
	},
}
