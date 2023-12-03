package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:               "httpushover",
	Short:             "httpushover - Automated login for HSP-Hamburg",
	Long:              `Automatically login to your HSP-Courses on the HSP-Hamburg Website.`,
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Run Root CMD")
	},
}

func init() {

	RootCmd.AddCommand(loginCmd)
	RootCmd.AddCommand(serverCmd)
}
