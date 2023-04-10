// Copyright (C) 2023, WEAVEDB LTD. All rights reserved.
// See the file LICENSE for licensing terms.

// "weavedb-cli" implements tokenvm client operation interface.
package main

import (
	"os"

	"github.com/ava-labs/hypersdk/examples/tokenvm/cmd/weavedb-cli/cmd"
	"github.com/ava-labs/hypersdk/utils"
)

func main() {
	if err := cmd.Execute(); err != nil {
		utils.Outf("{{red}}weavedb-cli exited with error:{{/}} %+v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
