// Copyright 2018 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

// Annotations from the CRI annotations package.
//
// These are vendor due to import conflicts.
const (
	sandboxLogDirAnnotation = "io.kubernetes.cri.sandbox-log-directory"
	// ContainerTypeAnnotation is they key that defines sandbox or container.
	ContainerTypeAnnotation = "io.kubernetes.cri.container-type"
	containerTypeSandbox    = "sandbox"
	// ContainerTypeContainer is the value for container.
	ContainerTypeContainer = "container"
)
