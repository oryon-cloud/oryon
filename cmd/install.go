// Copyright 2023 The Oryon Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/oryon-cloud/oryon/config"
	"github.com/oryon-cloud/oryon/logger"
	"github.com/oryon-cloud/oryon/pkm"
	"github.com/spf13/cobra"
)

// flags
var Force bool // force overwrite

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Oryon Application",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.Flags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.oryon.toml)")
		// check args
		if len(args) == 0 {
			logger.Logger.Error("Please specify the module name")
			os.Exit(1)
		}

		// read config file
		if configFile != "" {
			err := config.ParseConfig(configFile)
			if err != nil {
				logger.Logger.Error("Error reading config file", err)
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range args {
			logger.Logger.Info("Installing module " + name)
			m := pkm.PackageManager{Name: name}
			if err := m.Install(); err != nil {
				logger.Logger.Error("Error installing module", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	InstallCmd.Flags().BoolVarP(&Force, "force", "f", false, "force overwrite")
	config.Config.BindPFlag("force", InstallCmd.Flags().Lookup("force"))
}
