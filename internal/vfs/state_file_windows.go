// Copyright 2021 Google LLC
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

//go:build windows
// +build windows

package vfs

import (
	"fmt"
	"path/filepath"

	"golang.org/x/sys/windows"
)

func knownFolder(kind StateKind) (string, error) {
	switch kind {
	case Config:
		return windows.KnownFolderPath(windows.FOLDERID_LocalAppData, windows.KF_FLAG_CREATE)
	case SavedGames:
		return windows.KnownFolderPath(windows.FOLDERID_SavedGames, windows.KF_FLAG_CREATE)
	default:
		return "", fmt.Errorf("searched for unsupported state kind: %d", kind)
	}
}

func pathForReadRaw(kind StateKind, name string) ([]string, error) {
	path, err := pathForWrite(kind, name)
	return []string{path}, err
}

func pathForWriteRaw(kind StateKind, name string) (string, error) {
	root, err := knownFolder(kind)
	if err != nil {
		return "", err
	}
	return filepath.Join(root, gameName, name), nil
}
