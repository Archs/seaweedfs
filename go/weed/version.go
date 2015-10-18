package main

import (
	"fmt"
	"runtime"

	"github.com/chrislusf/seaweedfs/go/util"
)

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print DFS version",
	Long:      `Version prints the DFS version`,
}

func runVersion(cmd *Command, args []string) bool {
	if len(args) != 0 {
		cmd.Usage()
	}

	fmt.Printf("version %s %s %s\n", util.VERSION, runtime.GOOS, runtime.GOARCH)
	return true
}
