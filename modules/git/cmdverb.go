// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package git

import (
	"code.gitea.io/gitea/models/perm"
)

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
		CmdVerbLfsTransfer,
		CmdVerbAnnexShell:
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

// AnnexAccessMode returns the access mode required for a git-annex-shell subcommand.
func AnnexAccessMode(subVerb string) perm.AccessMode {
	switch subVerb {
	case "recvkey", "dropkey", "lockcontent", "unlockcontent":
		return perm.AccessModeWrite
	default: // configlist, inannex, sendkey, notifychanges, p2pstdio, etc.
		return perm.AccessModeRead
	}
}
