package main

import (
	"github.com/datianshi/concourse-cnb/config/pkg/env"
	"github.com/spf13/cobra"
)

var prefix, output string

func init() {
	envCmd.Flags().StringVarP(&prefix, "prefix", "p", "BUILD_ENV_", "prefix environment variable")
	envCmd.Flags().StringVarP(&output, "output", "o", "env", "output folder")
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Generate buildpack env",
	Run: func(cmd *cobra.Command, args []string) {
		if err := env.EnvBuildpack(prefix, output); err != nil {
			panic(err)
		}
	},
}
