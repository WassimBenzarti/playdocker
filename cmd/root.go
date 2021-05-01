package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "",
}

func Execute() {
	rootCmd.AddCommand(StartCmd)
	rootCmd.AddCommand(SetCookieCmd)
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}