// NOTE: This assumes you are building a cli and using cobra
package PLACEHOLDER	// CHANGE <name> to you package name 

import (
	"go.szostok.io/version/extension"
)

const (
	RepoOwner string = "PLACEHOLDER" // TODO: update this value
	RepoName  string = "PLACEHOLDER" // TODO: update this value
)

func init() {
	RootCmd.AddCommand(
		extension.NewVersionCobraCmd(
			extension.WithUpgradeNotice(RepoOwner, RepoName),
		),
	)
}