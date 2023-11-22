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

package config

import (
	"os"
	"path/filepath"

	"github.com/oryon-cloud/oryon/logger"
	"github.com/spf13/viper"
)

var Config = viper.New()

func ParseConfig(configFile string) error {
	Config.SetConfigFile(configFile)
	err := Config.ReadInConfig()
	if err != nil {
		logger.Logger.Error("Error reading config file", err)
		return err
	}
	return nil
}

func init() {
	Config.SetEnvPrefix("ORYON")
	Config.AutomaticEnv()
	// set default config

	hdir, err := os.UserHomeDir()
	if err != nil {
		logger.Logger.Error("Error getting home directory", err)
	}

	Config.SetDefault("addons_path", filepath.Join(hdir, ".oryon", "addons"))
	Config.SetDefault("npm_path", filepath.Join(hdir, ".oryon", "node_modules"))
	Config.SetDefault("cache_path", filepath.Join(hdir, ".oryon", "cache"))

}
