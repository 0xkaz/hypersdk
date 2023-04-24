// Copyright (C) 2023, WEAVEDB LTD. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/ulimit"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm"
	"github.com/ava-labs/hypersdk/examples/tokenvm/cmd/weavedbvm/version"
	"github.com/ava-labs/hypersdk/examples/tokenvm/controller"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:        "weavedbvm",
	Short:      "WeavedbVM agent",
	SuggestFor: []string{"weavedbvm"},
	RunE:       runFunc,
}

func init() {
	cobra.EnablePrefixMatching = true
}

func init() {
	rootCmd.AddCommand(
		version.NewCommand(),
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "weavedbvm failed %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func runFunc(*cobra.Command, []string) error {
	log.Printf("runFu")
	if err := ulimit.Set(ulimit.DefaultFDLimit, logging.NoLog{}); err != nil {
		return fmt.Errorf("%w: failed to set fd limit correctly", err)
	}
	return rpcchainvm.Serve(context.TODO(), controller.New())
}
