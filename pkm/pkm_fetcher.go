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
	"net/url"
	"os"
	"strings"

	"github.com/oryon-cloud/oryon/app"
)

type Fetcher interface {
	Fetch() (*app.ModuleManifest, error)
}

// check name is local path
func isLocalPath(name string) bool {
	if fileinfo, err := os.Stat(name); err == nil && fileinfo.IsDir() {
		return true
	}
	return false
}

// check name is remote path
func isRemotePath(name string) bool {
	if f, err := url.Parse(name); err != nil || f.Scheme == "" || f.Host == "" || f.Path == "" {
		return false
	}
	return true
}

// check name is git repository
func isGitRepository(name string) bool {
	// <protocol>://[<user>[:<password>]@]<hostname>[:<port>][:][/]<path>[#<commit-ish> | #semver:<semver>]
	// <protocol> is one of git, git+ssh, git+http, git+https, or git+file.
	return strings.HasPrefix(name, "git://") || strings.HasPrefix(name, "git+ssh://") || strings.HasPrefix(name, "git+http://") || strings.HasPrefix(name, "git+https://") || strings.HasPrefix(name, "git+file://")
}

// createFetcher creates a fetcher based on the nameStr
func NewFetcher(name string) (Fetcher, error) {
	if isLocalPath(name) {
		return newLocalFetcher(name)
	} else if isRemotePath(name) {
		return newRemoteFetcher(name)
	} else {
		return newRegistryFetcher(name)
	}
}
