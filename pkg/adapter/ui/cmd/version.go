package cmd

import (
	"fmt"
)

const versionFormat = "%s (build %s; commit %s)"

func init() {
	rootCmd.Version = fmt.Sprintf(versionFormat, Version, Commit, Date)
}
