// Copyright 2021 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package private

import (
	"net/http"

	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/private"
	"code.gitea.io/gitea/modules/setting"
	"code.gitea.io/gitea/modules/web"
	"code.gitea.io/gitea/services/context"
)

// SSHLog hook to response ssh log
func SSHLog(ctx *context.PrivateContext) {
	opts := web.GetForm(ctx).(*private.SSHLogOption)

	if opts.IsError {
		// Always log SSH errors so they are visible in the system journal even
		// when EnableSSHLog is false. This is important for diagnosing failures
		// in git-annex-shell and other SSH sub-commands.
		log.Error("ssh: %v", opts.Message)
		ctx.Status(http.StatusOK)
		return
	}

	if setting.Log.EnableSSHLog {
		log.Debug("ssh: %v", opts.Message)
	}
	ctx.Status(http.StatusOK)
}
