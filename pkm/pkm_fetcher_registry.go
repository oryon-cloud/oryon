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
	"errors"
	"net/http"
	"strings"

	"github.com/oryon-cloud/oryon/app"
)

type registryFetcher struct {
	*remoteFetcher        // inherit remote fetcher
	registryUrl    string // registry url, default: https://github.com/oryon-cloud/registry
}

// parseNameAndVersion parses the name and version of the module
func (f *registryFetcher) parseNameAndVersion() error {
	if strings.Contains(f.name, "@") {
		nameAndVersion := strings.Split(f.name, "@")
		f.name = nameAndVersion[0]
		f.version = nameAndVersion[1]
	} else {
		f.version = "latest"
	}
	return nil
}

// getRawUrl returns the raw url of the manifest.json
func (f *registryFetcher) getRegistryManifestUrl() (string, error) {
	if strings.Contains(f.registryUrl, "https://github.com") {
		ownerAndRepo := strings.Split(strings.TrimPrefix(f.registryUrl, "https://github.com/"), "/")
		owner := ownerAndRepo[0]
		repo := ownerAndRepo[1]
		return "https://raw.githubusercontent.com/" + owner + "/" + repo + "/main/addons/" + f.name + "/" + f.version + "/manifest.json", nil
	}
	return "", errors.New("unsupported registry url")
}

// getDownloadUrl returns the download url of the module
func (f *registryFetcher) getDownloadUrl() error {
	if err := f.parseNameAndVersion(); err != nil {
		return err
	}
	registryManifestUrl, err := f.getRegistryManifestUrl()
	if err != nil {
		return err
	}
	resp, err := http.Get(registryManifestUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var mmf app.ModuleManifest
	if err := json.NewDecoder(resp.Body).Decode(&mmf); err != nil {
		return err
	}

	if mmf.Tarball != "" {
		f.downloadUrl = mmf.Tarball
	} else if mmf.Repository != "" {
		f.downloadUrl = mmf.Repository + "/archive/refs/tags/" + mmf.Version + ".tar.gz"
	} else {
		return errors.New("invalid manifest.json")
	}

	return nil
}

// impl Fetcher interface
func (f *registryFetcher) Fetch() (*app.ModuleManifest, error) {
	if err := f.getDownloadUrl(); err != nil {
		return nil, err
	}

	mm, err := f.remoteFetcher.Fetch()
	if err != nil {
		return nil, err
	}
	return mm, nil
}

// newRegistryFetcher creates a registry fetcher
func newRegistryFetcher(name string) (Fetcher, error) {
	return &registryFetcher{
		remoteFetcher: &remoteFetcher{
			localFetcher: &localFetcher{
				name: name,
			},
		},
		registryUrl: "https://github.com/oryon-cloud/registry",
	}, nil
}
