// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"

	"github.com/spf13/cobra"

	// "github.com/ava-labs/hypersdk/examples/tokenvm/consts"
	// "github.com/ava-labs/hypersdk/examples/tokenvm/version"
	"github.com/ava-labs/hypersdk/examples/weavedbvm/consts"
	"github.com/ava-labs/hypersdk/examples/weavedbvm/version"
)

func init() {
	cobra.EnablePrefixMatching = true
}

// NewCommand implements "weavedbvm version" command.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prints out the verson",
		RunE:  versionFunc,
	}
	return cmd
}

func versionFunc(*cobra.Command, []string) error {
	fmt.Printf("%s@%s (%s)\n", consts.Name, version.Version, consts.ID)
	return nil
}
