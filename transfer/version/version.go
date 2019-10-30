package version

import "fmt"

var (
	// Plugin name
	Plugin = "transfer"

	// Image name
	Image = "wolfulus/transfer"

	// Version holds the complete version number (link time)
	Version = "latest"

	// FQDN holds the full image name with its tag
	FQDN = "wolfulus/transfer:latest"
)

func init() {
	FQDN = fmt.Sprintf("%s:%s", Image, Version)
}
