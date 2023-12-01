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

type Model struct {
	Domain      string       `json:"domain"`
	Module      string       `json:"module"`
	Name        string       `json:"name"`
	Decorators  []*Decorator `json:"decorators"`
	Extends     *Extends     `json:"extends"`
	Constructor *Method      `json:"constructor"`
	Properties  []*Property  `json:"properties"`
	Methods     []*Method    `json:"methods"`
	Source      string       `json:"source"`
}

type Extends struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
	Module string `json:"module"`
}

type Parameter struct {
	Name string `json:"name"`
	Type *Type  `json:"type"`
	// Types []string `json:"types"`
}

type Type struct {
	Name          string  `json:"name"`
	Types         []*Type `json:"types"`
	TypeName      string  `json:"typeName"`
	TypeArguments []*Type `json:"typeArguments"`
}

type Decorator struct {
	Name      string      `json:"name"`
	Arguments []*Argument `json:"arguments"`
}

type Argument struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

type Property struct {
	Name       string       `json:"name"`
	Type       *Type        `json:"type"`
	Decorators []*Decorator `json:"decorators"`
	Modifiers  *Modifiers   `json:"modifiers"`
	Setter     *Method      `json:"setter"`
	Getter     *Method      `json:"getter"`
}

type Method struct {
	Name       string       `json:"name"`
	ReturnType *Type        `json:"type"`
	Decorators []*Decorator `json:"decorators"`
	Parameters []*Parameter `json:"parameters"`
	Modifiers  *Modifiers   `json:"modifiers"`
}

type Modifiers struct {
	Public    bool `json:"public"`
	Private   bool `json:"private"`
	Protected bool `json:"protected"`
	Static    bool `json:"static"`
	Abstract  bool `json:"abstract"`
	Async     bool `json:"async"`
	Declare   bool `json:"declare"`
	Default   bool `json:"default"`
	Export    bool `json:"export"`
}
