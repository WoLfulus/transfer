package commands

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/wolfulus/transfer/transfer/service"
)

type installOptions struct {
	port     int
	address  string
	username string
	password string
	env      []string
	force    bool
}

func buildInstall() *cobra.Command {

	options := installOptions{}

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install transfer server in the current host",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInstall(options)
		},
	}

	flags := cmd.Flags()

	flags.IntVar(&options.port, "port", -1, "If set, exposes the service on the specified port")
	flags.StringVar(&options.address, "address", "0.0.0.0", "Address to bind the port if exposed")
	flags.StringVar(&options.username, "username", "", "The system username (used for internal push/pull)")
	flags.StringVar(&options.password, "password", "", "The system password (used for internal push/pull)")
	flags.StringArrayVar(&options.env, "env", []string{}, "Environment variables to put in the service container")

	flags.BoolVar(&options.force, "force", false, "Forces the installation")

	return cmd
}

func runInstall(options installOptions) error {
	svc, err := service.Get()
	if err != nil && err != service.ErrServiceNotFound {
		return err
	}

	if svc.Status == service.StatusUnknown {
		service.Error("Service status could not be determined. Aborting.")
		os.Exit(1)
		return nil
	}

	if svc.Status == service.StatusMultiple {
		service.Error("Impossible to install because multiple services are already running.")
		service.Error("Please remove all previous service containers first.")
		os.Exit(1)
		return nil
	}

	if svc.Status == service.StatusOK {
		if options.force == false {
			service.Error("There is another service installed. Please uninstall the previous service first or use --force to override the current service.")
			os.Exit(1)
			return nil
		}

		if svc.Managed == false {
			service.Error("The current service cannot be replaced because it wasn't created by transfer CLI.")
			os.Exit(1)
			return nil
		}

		svc.Uninstall()
	}

	if options.username == "" {
		options.username = "transfer"
	}

	if options.password == "" {
		options.password = generatePassword(16)
	}

	err = service.Install(options.username, options.password, options.port, options.address, options.env)
	if err != nil {
		service.Error("Service installation failed: %s", err)
		os.Exit(1)
	} else {
		service.Log("Tranasfer service installed")
	}

	return err
}

func generatePassword(n int) string {
	const letterBytes = "abcdefghkmnpqrstuvwxyzABCDEFGHKMNPQRSUVWXYZ123456789"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
