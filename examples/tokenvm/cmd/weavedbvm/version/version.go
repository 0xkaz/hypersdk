// Copyright (C) 2023, WEAVEDB LTD. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/ava-labs/hypersdk/examples/tokenvm/consts"
	"github.com/ava-labs/hypersdk/examples/tokenvm/version"
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
	fmt.Printf("%s", GetVersion())
	return nil
}
func GetVersion() string {
	return fmt.Sprintf("%s@%s (%s)\n", consts.Name, version.Version, consts.ID)
}
