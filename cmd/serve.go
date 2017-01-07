package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/rms1000watt/hello-world-go-grpc/src"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	address string
	logging bool
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Hello World server",
	Long: `Start the Hello World server

This is a gRPC based server. Use the pb/helloWorld.proto to understand
the RPC contract.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		config := src.Config{
			Address: address,
			Logging: logging,
		}

		log.Println("CONFIG", config)
		src.Serve(config)
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringVarP(&address, "address", "a", ":8081", "Address to listen on")
	serveCmd.PersistentFlags().BoolVarP(&logging, "logging", "l", false, "Enable logging")

	// Catch and add env vars to pflags
	// Courtesy of https://github.com/coreos/pkg/blob/master/flagutil/env.go
	serveCmd.PersistentFlags().VisitAll(func(f *pflag.Flag) {
		key := strings.ToUpper(strings.Replace(f.Name, "-", "_", -1))
		if val := os.Getenv(key); val != "" {
			if err := serveCmd.PersistentFlags().Set(f.Name, val); err != nil {
				log.Println("ERROR", err)
			}
		}
	})
}
