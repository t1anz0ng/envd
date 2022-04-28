// Copyright 2022 The MIDI Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ir

import (
	"github.com/moby/buildkit/client/llb"

	"github.com/tensorchord/MIDI/pkg/vscode"
)

// A Graph contains the state,
// such as its call stack and thread-local storage.
type Graph struct {
	OS       string
	Language string
	CUDA     *string
	CUDNN    *string

	UbuntuAPTSource *string
	PyPIMirror      *string

	BuiltinSystemPackages []string
	PyPIPackages          []string
	SystemPackages        []string
	VSCodePlugins         []vscode.Plugin

	Exec []llb.State
}