package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wassimbenzarti/pwd-go/lib/session_service"
)

var SetCookieCmd = &cobra.Command{
	Use:   "login {cookie to store}",
	Short: "Set cookie as credential",
	Long:  `Set cookie as credential`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sessionService, err := session_service.New()
		if err != nil {
			panic(err)
		}
		sessionService.SetCredentials(args[0])
		fmt.Printf("Credentials were saved in: %s", sessionService.GetConfigPath())
	},
}
