package cmd

import "github.com/spf13/cobra"

var bindingName string
var output string
var bindingType string

func init() {
	BindingsCmd.PersistentFlags().StringVarP(&bindingName, "binding-name", "n", "", "binding name")
	BindingsCmd.MarkFlagRequired("bindingName")
	BindingsCmd.PersistentFlags().StringVarP(&bindingType, "binding-type", "t", "", "binding type")
	BindingsCmd.MarkFlagRequired("bindingType")
	BindingsCmd.PersistentFlags().StringVarP(&output, "output", "o", "bindings", "output folder")
	BindingsCmd.AddCommand(mavenSettingsCmd)
	BindingsCmd.AddCommand(datadogCmd)
}

var BindingsCmd = &cobra.Command{
	Use:   "bindings",
	Short: "Generate bindings",
}
