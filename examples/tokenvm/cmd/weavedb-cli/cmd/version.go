package cmd

import (
	// "github.com/fatih/color"
	// "github.com/ava-labs/hypersdk/examples/tokenvm/cmd/tokenvm/version"
	"github.com/ava-labs/hypersdk/examples/tokenvm/version"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	RunE: func(*cobra.Command, []string) error {
		color.Green("Version: %v", version.GetVersion())
		return nil
	},
}
