// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package git

const (
	CmdVerbUploadPack      = "git-upload-pack"
	CmdVerbUploadArchive   = "git-upload-archive"
	CmdVerbReceivePack     = "git-receive-pack"
	CmdVerbLfsAuthenticate = "git-lfs-authenticate"
	CmdVerbLfsTransfer     = "git-lfs-transfer"
	CmdVerbAnnexShell      = "git-annex-shell"

	CmdSubVerbLfsUpload   = "upload"
	CmdSubVerbLfsDownload = "download"
)

func IsAllowedVerbForServe(verb string) bool {
	switch verb {
	case CmdVerbUploadPack,
		CmdVerbUploadArchive,
		CmdVerbReceivePack,
		CmdVerbLfsAuthenticate,
		CmdVerbLfsTransfer:
		return true
	}
	return false
}

func IsAllowedVerbForServeLfs(verb string) bool {
	switch verb {
	case CmdVerbLfsAuthenticate,
		CmdVerbLfsTransfer:
		return true
	}
	return false
}

// IsAllowedVerbForServeAnnex checks if the verb is a git-annex command.
func IsAllowedVerbForServeAnnex(verb string) bool {
	return verb == CmdVerbAnnexShell
}

