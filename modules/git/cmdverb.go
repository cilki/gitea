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

// AnnexShellExecutable is the resolved path to git-annex-shell. It is
// initialized by InitSimple using setting.Git.AnnexShellPath and must be
// used instead of a plain "git-annex-shell" string so that gitea serv
// (spawned by sshd with a restricted PATH) can find the binary.
var AnnexShellExecutable = CmdVerbAnnexShell

