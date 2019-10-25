package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Short: "Direct image deployments",
    Use:   "transfer",
}

func Execute() error {
    return rootCmd.Execute()
}