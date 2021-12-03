package cmd

import (
	"github.com/datianshi/concourse-cnb/config/pkg/bindings"
	"github.com/datianshi/concourse-cnb/config/pkg/bindings/datadog"
	"github.com/spf13/cobra"
)

var apiKey string

func init() {
	datadogCmd.Flags().StringVarP(&apiKey, "api-key", "a", "", "datadog api key")
	datadogCmd.MarkFlagRequired("api-key")
}

var datadogCmd = &cobra.Command{
	Use:   "datadog",
	Short: "Generate datadog binding",
	Run: func(cmd *cobra.Command, args []string) {
		s := datadog.NewDataDog(apiKey)
		binding := bindings.NewBinding(bindingName, "DatadogTrace", output, s)
		if err := binding.CreateBinding(); err != nil {
			panic(err)
		}
	},
}
