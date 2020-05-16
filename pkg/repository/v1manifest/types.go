// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package v1manifest

// Manifest representation for ser/de.
type Manifest struct {
	// Signatures value
	Signatures []signature `json:"signatures"`
	// Signed value; any value here must have the SignedBase base.
	Signed ValidManifest `json:"signed"`
}

type signature struct {
	KeyID string `json:"keyid"`
	Sig   string `json:"sig"`
}

// SignedBase represents parts of a manifest's signed value which are shared by all manifests.
type SignedBase struct {
	Ty          string `json:"_type"`
	SpecVersion string `json:"spec_version"`
	Expires     string `json:"expires"`
	// 0 => no version specified
	Version uint `json:"version"`
}

// ValidManifest is a manifest which includes SignedBase and can be validated.
type ValidManifest interface {
	isValid() error
	// Base returns this manifest's SignedBase which is values common to all manifests.
	Base() *SignedBase
	// Filename returns the unversioned name that the manifest should be saved as based on its Go type.
	Filename() string
}

// Root manifest.
type Root struct {
	SignedBase
	Roles map[string]*Role `json:"roles"`
}

// Role object.
type Role struct {
	URL       string              `json:"url"`
	Keys      map[string]*KeyInfo `json:"keys"`
	Threshold uint                `json:"threshold"`
}

// KeyInfo is the manifest structure of a single key
type KeyInfo struct {
	Algorithms []string          `json:"keyid_hash_algorithms"`
	Type       string            `json:"keytype"`
	Value      map[string]string `json:"keyval"`
	Scheme     string            `json:"scheme"`
}

// ComponentItem object
type ComponentItem struct {
	Name      string `json:"name"`
	Yanked    bool   `json:"yanked"`
	Owner     string `json:"owner"`
	URL       string `json:"url"`
	Length    int64  `json:"length"`
	Threshold int    `json:"threshold"`
}

// Index manifest.
type Index struct {
	SignedBase
	Owners            map[string]Owner         `json:"owners"`
	Components        map[string]ComponentItem `json:"components"`
	DefaultComponents []string                 `json:"default_components"`
}

// Owner object.
type Owner struct {
	Name string              `json:"name"`
	Keys map[string]*KeyInfo `json:"keys"`
}

// VersionItem object
type VersionItem struct {
	Yanked       bool              `json:"yanked"`
	URL          string            `json:"url"`
	Hashes       map[string]string `json:"hashes"`
	Entry        string            `json:"entry"`
	Released     string            `json:"released"`
	Length       int64             `json:"length"`
	Dependencies []string          `json:"dependencies"`
}

// Component manifest.
type Component struct {
	SignedBase
	Name        string                            `json:"name"`
	Description string                            `json:"description"`
	Platforms   map[string]map[string]VersionItem `json:"platforms"`
}

// Snapshot manifest.
type Snapshot struct {
	SignedBase
	Meta map[string]FileVersion `json:"meta"`
}

// Timestamp manifest.
type Timestamp struct {
	SignedBase
	Meta map[string]FileHash `json:"meta"`
}

// FileHash is the hashes and length of a file.
type FileHash struct {
	Hashes map[string]string `json:"hashes"`
	Length uint              `json:"length"`
}

// FileVersion is just a version number.
type FileVersion struct {
	Version uint `json:"version"`
}

// Boilerplate implementations of ValidManifest.

// Base implements ValidManifest
func (manifest *Root) Base() *SignedBase {
	return &manifest.SignedBase
}

// Base implements ValidManifest
func (manifest *Index) Base() *SignedBase {
	return &manifest.SignedBase
}

// Base implements ValidManifest
func (manifest *Component) Base() *SignedBase {
	return &manifest.SignedBase
}

// Base implements ValidManifest
func (manifest *Snapshot) Base() *SignedBase {
	return &manifest.SignedBase
}

// Base implements ValidManifest
func (manifest *Timestamp) Base() *SignedBase {
	return &manifest.SignedBase
}

// Filename implements ValidManifest
func (manifest *Root) Filename() string {
	return ManifestsConfig[ManifestTypeRoot].Filename
}

// Filename implements ValidManifest
func (manifest *Index) Filename() string {
	return ManifestsConfig[ManifestTypeIndex].Filename
}

// Filename implements ValidManifest
func (manifest *Component) Filename() string {
	panic("Unreachable")
}

// Filename implements ValidManifest
func (manifest *Snapshot) Filename() string {
	return ManifestsConfig[ManifestTypeSnapshot].Filename
}

// Filename implements ValidManifest
func (manifest *Timestamp) Filename() string {
	return ManifestsConfig[ManifestTypeTimestamp].Filename
}