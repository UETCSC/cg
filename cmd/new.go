package cmd

import (
	"cg/pkg/cmdutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newCmd)
	newCmd.AddCommand(WizardCmd)
	newCmd.AddCommand(FileCmd)
}

var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a new challenge",
	Long:    `Create a new challenge`,
	Aliases: []string{"n"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var WizardCmd = &cobra.Command{
	Use:     "wizard",
	Short:   "Create a new challenge using the wizard",
	Long:    `Create a new challenge using the wizard`,
	Aliases: []string{"w"},
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.Wizard()
	},
}

var FileCmd = &cobra.Command{
	Use:     "file",
	Short:   "Create a new challenge from predefined file",
	Long:    `Create a new challenge from predefined file`,
	Aliases: []string{"f"},
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.File()
	},
}
