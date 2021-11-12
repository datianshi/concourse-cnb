package main

import (
	"os"

	"github.com/datianshi/concourse-cnb/config/pkg/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "buildpack",
	Short: "buildpack config",
}

func init() {
	rootCmd.AddCommand(envCmd)
	rootCmd.AddCommand(cmd.BindingsCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
