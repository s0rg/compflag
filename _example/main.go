package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/s0rg/compflag"
)

var (
	fWorkers    = flag.Int("workers", 6, "number of workers")
	fSilent     = flag.Bool("silent", false, "suppress messages in stderr")
	fDirsPolicy = flag.String("dirs", "show", "policy for non-resource urls: show / hide / only")
	fProxyAuth  = flag.String("proxy-auth", "", "credentials for proxy: user:password")
	fDelay      = flag.Duration("delay", 100, "per-request delay (0 - disable)")
)

func main() {
	// option 1
	if compflag.Complete() {
		os.Exit(0)
	}

	// option 2
	// compflag.Var("complete")

	flag.Parse()

	fmt.Printf("Running '%s' with parameters:\n", os.Args[0])
}
