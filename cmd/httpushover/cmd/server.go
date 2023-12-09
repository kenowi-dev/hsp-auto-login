package cmd

import (
	"github.com/kenowi-dev/hsp-auto-login/pkg/api"
	"github.com/spf13/cobra"
)

type ServerConfig struct {
	port int
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  "Server that handles automatic hsp logins",
	Args:  cobra.NoArgs,
	Run:   runServer,
}
var serverConfig ServerConfig

// init sets up CLI flags for the version command.
func init() {
	serverCmd.Flags().IntVarP(&serverConfig.port, "port", "p", 8080, "The url for the login (default 8080)")
	//_ = serverCmd.MarkFlagRequired("pushoverKey")
}

func runServer(_ *cobra.Command, _ []string) {

	a := api.New()
	a.SetupAndRun(serverConfig.port)
}
