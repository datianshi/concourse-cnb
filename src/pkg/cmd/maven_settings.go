package cmd

import (
	"github.com/datianshi/concourse-cnb/config/pkg/bindings"
	"github.com/datianshi/concourse-cnb/config/pkg/bindings/maven"
	"github.com/spf13/cobra"
)

var username string
var password string
var repo string

func init() {
	mavenSettingsCmd.Flags().StringVarP(&username, "username", "u", "", "maven repo username")
	mavenSettingsCmd.MarkFlagRequired("username")
	mavenSettingsCmd.Flags().StringVarP(&password, "password", "p", "", "maven repo password")
	mavenSettingsCmd.MarkFlagRequired("password")
	mavenSettingsCmd.Flags().StringVarP(&repo, "repo", "r", "", "maven repo url")
	mavenSettingsCmd.MarkFlagRequired("repo")
}

var mavenSettingsCmd = &cobra.Command{
	Use:   "maven-settings",
	Short: "Generate maven settings",
	Run: func(cmd *cobra.Command, args []string) {
		s := maven.NewSettings(username, password, repo)
		binding := bindings.NewBinding(bindingName, "maven", output, s)
		if err := binding.CreateBinding(); err != nil {
			panic(err)
		}
	},
}
