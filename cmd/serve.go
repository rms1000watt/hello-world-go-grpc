// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
