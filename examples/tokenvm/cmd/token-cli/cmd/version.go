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
		color.Green("version %v", version.GetVersion())
		return nil
		// return ErrMissingSubcommand
	},
}

// var genVersionCmd = &cobra.Command{
// 	Use:   "version",
// 	Short: "version",
// 	PreRunE: func(cmd *cobra.Command, args []string) error {
// 		if len(args) != 1 {
// 			return ErrInvalidArgs
// 		}
// 		return nil
// 	},
// 	RunE: func(_ *cobra.Command, args []string) error {
// 		log.Printf("args: %v", args)
// 		// g := genesis.Default()
// 		// if minUnitPrice >= 0 {
// 		// 	g.MinUnitPrice = uint64(minUnitPrice)
// 		// }

// 		// a, err := os.ReadFile(args[0])
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// allocs := []*genesis.CustomAllocation{}
// 		// if err := json.Unmarshal(a, &allocs); err != nil {
// 		// 	return err
// 		// }
// 		// g.CustomAllocation = allocs

// 		// b, err := json.Marshal(g)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// if err := os.WriteFile(genesisFile, b, fsModeWrite); err != nil {
// 		// 	return err
// 		// }
// 		// color.Green("created genesis and saved to %s", genesisFile)
// 		color.Green("version")
// 		return nil
// 	},
// }
