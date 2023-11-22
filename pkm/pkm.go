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
	"github.com/oryon-cloud/oryon/app"
)

type PackageManager struct {
	Name  string // name or url of module
	Force bool   // force overwrite
}

// module install
func (pm *PackageManager) Install() error {
	fetch, err := NewFetcher(pm.Name)
	if err != nil {
		return err
	}

	mm, err := fetch.Fetch()
	if err != nil {
		return err
	}

	module := app.Module{
		Manifest: mm,
	}
	if err := module.Install(); err != nil {
		return err
	}

	return nil
}

// module uninstall
func (pm *PackageManager) Uninstall() error {
	return nil
}

// module update
func (pm *PackageManager) Update() error {
	return nil
}
