// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"

	"github.com/ava-labs/avalanchego/version"
	"github.com/ava-labs/hypersdk/examples/tokenvm/consts"
)

var Version = &version.Semantic{
	Major: 0,
	Minor: 0,
	Patch: 1,
}

func GetVersion() string {
	return fmt.Sprintf("%s@%s (%s)\n", consts.Name, Version, consts.ID)
}
