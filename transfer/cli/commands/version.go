package commands

import (
    "fmt"

    "github.com/wolfulus/transfer/transfer/version"

	"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(versionCmd)
}

func versionRun(cmd *cobra.Command, args []string) {
    fmt.Printf("Transfer Version: %s\n", version.Version)
}

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version information of transfer plugin",
    Long:  "Print the version information of transfer plugin",
    Run:   versionRun,
}
