package cmd

import (
	"cg/pkg/cmdutil"
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(DockerCmd)
	DockerCmd.AddCommand(AutoCmd)
	DockerCmd.AddCommand(BuildCmd)
	DockerCmd.AddCommand(RunCmd)
	DockerCmd.AddCommand(StopCmd)
	DockerCmd.AddCommand(LogCmd)
	DockerCmd.AddCommand(SaveCmd)
}

var Cyan = color.FgCyan.Render
var Red = color.FgRed.Render

func CheckDockerCompose() bool {
	_, err := os.Stat("docker-compose.yml")
	if err == nil {
		fmt.Println(Cyan("Detected docker-compose.yml"))
	} else if os.IsNotExist(err) {
		fmt.Println(Red("No docker-compose.yml detected"))
	}
	return err == nil
}

var DockerCmd = &cobra.Command{
	Use:     "docker",
	Short:   "docker operations",
	Long:    `docker operations`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println()
		CheckDockerCompose()
		os.Exit(0)
	},
}

var AutoCmd = &cobra.Command{
	Use:     "auto",
	Short:   "Automate the workflow Stop -> Build -> Run",
	Long:    `Automate the workflow Stop -> Build -> Run`,
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start testing")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Auto()
	},
}

var BuildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Build image",
	Long:    `Build image`,
	Aliases: []string{"b"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building the image")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Build()
	},
}

var RunCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run image",
	Long:    `Run umage`,
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running the image")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Run()
	},
}

var StopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "Stop the container",
	Long:    `Stop the container`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Stop()
	},
}

var LogCmd = &cobra.Command{
	Use:   "log",
	Short: "View container logs",
	Long:  `View container logs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("View container logs")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Log()
	},
}

var SaveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the image to a tarball.",
	Long:  `Save the image to a tarball.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Save the image to a tarball.")
		if !CheckDockerCompose() {
			os.Exit(0)
		}
		cmdutil.Save()
	},
}
