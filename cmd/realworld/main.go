package main

import (
	"fmt"
	"os"

	"github.com/minchao/go-realworld/cmd/realworld/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
