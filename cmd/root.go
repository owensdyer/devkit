// Package cmd is the root command-line handler which passes arguments onto child commands.
package cmd

import (
	"devkit/cmd/docker"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the first command to initiate the package.
var RootCmd = &cobra.Command{
	Use:   "dk",
	Short: "Development Utility Kit",
	Long:  `Developer Kit (dk) is a small command-line tool used to remove all repetitive tasks that comes with development`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	RootCmd.AddCommand(docker.DockerCmd)

	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
