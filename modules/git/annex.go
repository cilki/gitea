// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package git

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// IsAnnexSymlink returns true if the symlink target points to a git-annex object.
func IsAnnexSymlink(symlinkTarget string) bool {
	return strings.Contains(symlinkTarget, ".git/annex/objects/")
}

// ResolveAnnexObjectPath resolves a git-annex symlink target to an absolute
// filesystem path within a bare repository. Bare repos store annex objects at
// <repoPath>/annex/objects/... rather than <repoPath>/.git/annex/objects/...
func ResolveAnnexObjectPath(repoPath, symlinkTarget string) (string, error) {
	// Strip the ".git/" prefix since bare repos don't have a .git directory
	rel := strings.TrimPrefix(symlinkTarget, ".git/")
	// Also handle targets with leading "../" components pointing to .git/annex
	for strings.HasPrefix(rel, "../") {
		rel = strings.TrimPrefix(rel, "../")
		if strings.HasPrefix(rel, ".git/annex/") {
			rel = strings.TrimPrefix(rel, ".git/")
			break
		}
	}
	absPath := filepath.Join(repoPath, rel)
	if _, err := os.Stat(absPath); err != nil {
		return "", fmt.Errorf("annex object not found: %w", err)
	}
	return absPath, nil
}

// OpenAnnexObject opens the annex object referenced by a symlink target and
// returns a ReadCloser, the file size, and any error.
func OpenAnnexObject(repoPath, symlinkTarget string) (io.ReadCloser, int64, error) {
	absPath, err := ResolveAnnexObjectPath(repoPath, symlinkTarget)
	if err != nil {
		return nil, 0, err
	}
	f, err := os.Open(absPath)
	if err != nil {
		return nil, 0, err
	}
	info, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, 0, err
	}
	return f, info.Size(), nil
}
