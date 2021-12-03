package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Fantom-foundation/go-opera/cmd/opera/launcher"
)

func main() {
	// increase max goroutines
	oldMax := runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	fmt.Println("Old max procs:", oldMax, "new max procs:", runtime.NumCPU() * 2)
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
