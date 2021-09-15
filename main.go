package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/takuzoo3868/go-msfdb/commands"
)

func main() {
	if envArgs := os.Getenv("GO_MSFDB_ARGS"); 0 < len(envArgs) {
		commands.RootCmd.SetArgs(strings.Fields(envArgs))
	}

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
