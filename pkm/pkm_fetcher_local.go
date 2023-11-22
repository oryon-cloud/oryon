// Copyright 2023 The Oryon Authors
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

package pkm

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/oryon-cloud/oryon/app"
	"github.com/oryon-cloud/oryon/config"
	"github.com/oryon-cloud/oryon/logger"
	cp "github.com/otiai10/copy"
)

type localFetcher struct {
	name           string
	version        string
	localPath      string
	moduleManifest *app.ModuleManifest
}

// copyModule copies the module to the addons_path
func (f *localFetcher) copyModule() error {
	destPath := filepath.Join(config.Config.GetString("addons_path"), f.name)

	// skip if already exists
	if f.localPath == destPath {
		logger.Logger.Info("Module already exists, skip copying")
		return nil
	}

	// if destPath exists and force is false, return error
	if _, err := os.Stat(destPath); err == nil && !config.Config.GetBool("force") {
		return os.ErrExist
	}

	// copy to destPath
	if err := cp.Copy(f.localPath, destPath, cp.Options{
		Skip: func(srcinfo os.FileInfo, src, dest string) (bool, error) {
			if srcinfo.IsDir() && srcinfo.Name() == ".git" {
				return true, nil
			}
			return false, nil
		},
	}); err != nil {
		return err
	}

	f.moduleManifest.Path = destPath

	return nil
}

// parseManifest parses the manifest.json and set the name and version
func (f *localFetcher) parseManifest() error {
	manifestFile, err := os.Open(filepath.Join(f.localPath, "manifest.json"))
	if err != nil {
		return err
	}
	defer manifestFile.Close()

	// parse manifest.json
	var mmf app.ModuleManifest
	if err := json.NewDecoder(manifestFile).Decode(&mmf); err != nil {
		return err
	}
	f.name = mmf.Name
	f.version = mmf.Version
	f.moduleManifest = &mmf

	return nil
}

// impl Fetcher interface
func (f *localFetcher) Fetch() (*app.ModuleManifest, error) {
	if err := f.parseManifest(); err != nil {
		return nil, err
	}
	if err := f.copyModule(); err != nil {
		return nil, err
	}
	return f.moduleManifest, nil
}

// newLocalFetcher creates a local fetcher
func newLocalFetcher(name string) (Fetcher, error) {
	if strings.HasPrefix(name, "file://") {
		name = strings.TrimPrefix(name, "file://")
	}

	// check name is directory
	if fileinfo, err := os.Stat(name); err != nil || !fileinfo.IsDir() {
		return nil, err
	}

	return &localFetcher{
		localPath: name,
	}, nil
}
