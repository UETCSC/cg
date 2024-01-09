package cmd

import (
	"cg/pkg/cmdutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(ConfigCmd)
	ConfigCmd.AddCommand(ConfigSetCmd)
	ConfigCmd.AddCommand(ConfigGetCmd)
	ConfigCmd.AddCommand(ConfigCleanCmd)
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure cg settings file.",
	Long:  "Configure cg settings file.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

var ConfigSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set current config.",
	Long:  "Set current config.",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigSet()
	},
}

var ConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get current config.",
	Long: "Get current config.",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigGet()
	},
}

var ConfigCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean current config.",
	Long:  "Clean current config.",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.ConfigClean()
	},
}
