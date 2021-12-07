package cmd

import (
	"github.com/datianshi/concourse-cnb/config/pkg/bindings"
	"github.com/spf13/cobra"
)

func init() {
}

var datadogCmd = &cobra.Command{
	Use:   "datadog",
	Short: "Generate datadog binding",
	Run: func(cmd *cobra.Command, args []string) {
		binding := bindings.NewBinding(bindingName, "DatadogTrace", output, &bindings.EmptyContent{})
		if err := binding.CreateBinding(); err != nil {
			panic(err)
		}
	},
}
