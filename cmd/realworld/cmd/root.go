package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version = "undefined"
	Date    = "undefined"
	Commit  = "undefined"
)

var (
	rootCmd = &cobra.Command{
		Use:   "realworld",
		Short: "the realworld example app",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
