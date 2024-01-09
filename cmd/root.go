package cmd

import (
	"os"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cg",
	Short: "cg - generate your CTF challenges with quality.",
	Long: `cg - generate your CTF challenges with quality. From uetctf with <3.`,
}


func Execute() {
	cc.Init(&cc.Config{
		RootCmd:         RootCmd,
		Headings:        cc.HiGreen + cc.Underline,
		Commands:        cc.Cyan + cc.Bold,
		Example:         cc.Italic,
		ExecName:        cc.Bold,
		Flags:           cc.Cyan + cc.Bold,
		NoExtraNewlines: true,
		NoBottomNewline: true,
	})
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
