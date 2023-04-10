// Copyright (C) 2023, WEAVEDB LTD. All rights reserved.
// See the file LICENSE for licensing terms.

// "weavedb-cli" implements weavedbvm client operation interface.
package cmd

import (
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/spf13/cobra"
)

const (
	requestTimeout  = 30 * time.Second
	fsModeWrite     = 0o600
	defaultDatabase = ".weavedb-cli"
	defaultGenesis  = "genesis.json"
)

var (
	dbPath string
	db     database.Database

	genesisFile     string
	minUnitPrice    int64
	hideTxs         bool
	randomRecipient bool
	maxTxBacklog    int

	rootCmd = &cobra.Command{
		Use:        "weavedb-cli",
		Short:      "WeavedbVM CLI",
		SuggestFor: []string{"weavedb-cli", "weavedbcli"},
	}
)

func init() {
	cobra.EnablePrefixMatching = true
	// rootCmd.AddCommand(
	// 	genesisCmd,
	// 	versionCmd,
	// 	keyCmd,
	// 	chainCmd,
	// 	actionCmd,
	// 	spamCmd,
	// )
	// rootCmd.PersistentFlags().StringVar(
	// 	&dbPath,
	// 	"database",
	// 	defaultDatabase,
	// 	"path to database (will create it missing)",
	// )
	// rootCmd.PersistentPreRunE = func(*cobra.Command, []string) error {
	// 	utils.Outf("{{yellow}}database:{{/}} %s\n", dbPath)
	// 	var err error
	// 	db, err = pebble.New(dbPath, pebble.NewDefaultConfig())
	// 	return err
	// }
	// rootCmd.PersistentPostRunE = func(*cobra.Command, []string) error {
	// 	return CloseDatabase()
	// }
	rootCmd.SilenceErrors = true

	// //// //// // genesis
	// genesisCmd.AddCommand(
	// 	genGenesisCmd,
	// )
	// genGenesisCmd.PersistentFlags().StringVar(
	// 	&genesisFile,
	// 	"genesis-file",
	// 	defaultGenesis,
	// 	"genesis file path",
	// )
	// genGenesisCmd.PersistentFlags().Int64Var(
	// 	&minUnitPrice,
	// 	"min-unit-price",
	// 	-1,
	// 	"minimum price",
	// )

	// //// //// // key
	// keyCmd.AddCommand(
	// 	genKeyCmd,
	// 	importKeyCmd,
	// 	setKeyCmd,
	// 	balanceKeyCmd,
	// )

	// //// //// //chain
	// watchChainCmd.PersistentFlags().BoolVar(
	// 	&hideTxs,
	// 	"hide-txs",
	// 	false,
	// 	"hide txs",
	// )
	// chainCmd.AddCommand(
	// 	importChainCmd,
	// 	importANRChainCmd,
	// 	setChainCmd,
	// 	chainInfoCmd,
	// 	watchChainCmd,
	// )

	// //// //// // actions
	// actionCmd.AddCommand(
	// 	transferCmd,

	// 	createAssetCmd,
	// 	mintAssetCmd,
	// 	// burnAssetCmd,
	// 	// modifyAssetCmd,

	// 	createOrderCmd,
	// 	fillOrderCmd,
	// 	closeOrderCmd,

	// 	importAssetCmd,
	// 	exportAssetCmd,
	// )

	// //// // spam
	// runSpamCmd.PersistentFlags().BoolVar(
	// 	&randomRecipient,
	// 	"random-recipient",
	// 	false,
	// 	"random recipient",
	// )
	// runSpamCmd.PersistentFlags().IntVar(
	// 	&maxTxBacklog,
	// 	"max-tx-backlog",
	// 	72_000,
	// 	"max tx backlog",
	// )
	// spamCmd.AddCommand(
	// 	runSpamCmd,
	// )
}

func Execute() error {
	return rootCmd.Execute()
}
