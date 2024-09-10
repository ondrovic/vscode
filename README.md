![GitHub License](https://img.shields.io/github/license/ondrovic/vscode)

# vscode
shared vscode settings

# usage
This is reliant on the [Repo Project Stubber](https://github.com/ondrovic/repo-project-stubber)


#### Used Extensions:

[Commands](https://marketplace.visualstudio.com/items?itemName=fabiospampinato.vscode-commands)

[Todo+](https://marketplace.visualstudio.com/items?itemName=fabiospampinato.vscode-todo-plus)


##### GO

[goreleaser](https://github.com/goreleaser/goreleaser)
[version](https://github.com/mszostok/version)

`version.go`
```go
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

```