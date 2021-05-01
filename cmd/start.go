package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wassimbenzarti/pwd-go/lib"
	"github.com/wassimbenzarti/pwd-go/lib/session_service"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a PWD session",
	Long:  `Start a session of play with docker`,
	Run: func(cmd *cobra.Command, args []string) {
		sessionService, err := session_service.New()
		if err != nil {
			panic(err)
		}

		if !sessionService.HasCredentials() {
			fmt.Println("You need to login first")
			return
		}

		sshHost := sessionService.GetSSHHost()
		fmt.Printf("You can connect to play with docker using this command:\n\tssh %s\n", sshHost)

		lib.LaunchSSHTunnel(sshHost)
	},
}
