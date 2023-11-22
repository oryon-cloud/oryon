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

package app

type ModuleManifest struct {
	Name                 string                       `json:"name"`                 // name, eg: foo
	Version              string                       `json:"version"`              // version, eg: v0.0.1
	Tarball              string                       `json:"tarball"`              // download tarball url, eg:
	Summary              string                       `json:"summary"`              // summary, eg: foo module
	Description          string                       `json:"description"`          // description, eg: foo module
	Domain               string                       `json:"domain"`               // domain name, eg: foo.oryon.cloud
	Main                 string                       `json:"main"`                 // main, eg: index.ts
	Depends              []string                     `json:"depends"`              // depends, eg: ["foo", "bar@v0.0.1"]
	ExternalDependencies map[string]map[string]string `json:"externalDependencies"` // external dependencies, eg: {"node_module": {"foo": "1.0.0"}, "binary": {"git": "0.12.28"}}
	Author               string                       `json:"author"`               // author, eg: Oryon team
	License              string                       `json:"license"`              // license, eg: Apache 2.0
	Homepage             string                       `json:"homepage"`             // homepage, eg: https://github.com/oryon-cloud
	Repository           string                       `json:"repository"`           // repository, eg: https://github.com/oryon-cloud/foo.git
	Path                 string                       // path, eg: /home/oryon/.oryon/addons/foo
}
