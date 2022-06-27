package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/vulsio/go-msfdb/commands"
)

func main() {
	if envArgs := os.Getenv("GO_MSFDB_ARGS"); 0 < len(envArgs) {
		commands.RootCmd.SetArgs(strings.Fields(envArgs))
	}

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
