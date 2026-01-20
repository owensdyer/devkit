// Package docker handles commands in relation to docker.
package docker

import (
	"devkit/cmd/docker/compose"

	"github.com/spf13/cobra"
)

// DockerCmd handles the primary command.
var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker related commands",
	Long:  `Docker command contains all commands that are relevant to docker.`,
}

func init() {
	DockerCmd.AddCommand(compose.ComposeCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dockerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dockerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
