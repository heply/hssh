/*
Copyright © 2021 Heply SRL <hello@heply.it>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"hssh/controllers"
	"os"

	"github.com/spf13/cobra"
)

// Version of the app provided
// in build phase
var Version string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "hssh",
	Short:   "A CLI to easily sync, list, search and connect to SSH hosts",
	Version: Version,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isInitCommand() bool {
	for _, arg := range os.Args {
		if arg == "init" || arg == "i" {
			return true
		}
	}
	return false
}

func init() {
	cobra.OnInitialize(func() {
		if isInitCommand() == false {
			controllers.Init(false)
		}
	})
}
