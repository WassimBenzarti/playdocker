package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "new",
    Short: "Launch a PWD session",
    Long: `Launch a session`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Hugo running")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}