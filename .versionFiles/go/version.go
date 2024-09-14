// NOTE: This assumes you are building a cli and using cobra
package PLACEHOLDER	// CHANGE <name> to you package name 

import (
	"go.szostok.io/version/extension"
)

const (
	// RepoOwner is the owner of the GitHub repository.
	RepoOwner string = "PLACEHOLDER" // TODO: update this value
	// RepoName is the name of the GitHub repository.
	RepoName  string = "PLACEHOLDER" // TODO: update this value
)

// init adds a new version command to RootCmd. The command includes an upgrade notice
// based on the specified GitHub repository owner and name.
func init() {
	RootCmd.AddCommand(
		extension.NewVersionCobraCmd(
			extension.WithUpgradeNotice(RepoOwner, RepoName),
		),
	)
}