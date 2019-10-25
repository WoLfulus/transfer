package cli

import (
	"fmt"
    "os"
    "encoding/json"
    "github.com/wolfulus/transfer/transfer/version"
    "github.com/wolfulus/transfer/transfer/cli/commands"
)

// Execute the command line application
func Execute() {
    if (len(os.Args) >= 2 && os.Args[1] == "docker-cli-plugin-metadata") {
        metadata, err := json.Marshal(Metadata{
            SchemaVersion: "0.1.0",
            Vendor: "WoLfulus.com",
            Version: version.Version,
            Experimental: false,
        })

        if err != nil {
            fmt.Println("Error printing plugin metadata")
            os.Exit(1)
        }

        fmt.Println(string(metadata))
        os.Exit(0)
    }

    if err := commands.Execute(); err != nil {
        os.Exit(1)
    }

    os.Exit(0)
}
