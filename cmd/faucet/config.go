package main

import (
	"flag"
	"strconv"

	"github.com/tendermint/starport/starport/pkg/cosmosfaucet"

	"github.com/tendermint/faucet/internal/environ"
)

const (
	commaSeparator = ","
)

var (
	port            int
	keyringBackend  string
	sdkVersion      string
	keyName         string
	keyMnemonic     string
	keyringPassword string
	appCli          string
	defaultDenoms   string
	creditAmount    string
	maxCredit       string
	nodeAddress     string
	legacySendCmd   bool
	coinType        string
	home            string
)

func init() {
	flag.IntVar(&port, "port",
		environ.GetInt("PORT", 8000),
		"tcp port where faucet will be listening for requests",
	)
	flag.StringVar(&keyringBackend, "keyring-backend",
		environ.GetString("KEYRING_BACKEND", ""),
		"keyring backend to be used",
	)
	flag.StringVar(&sdkVersion, "sdk-version",
		environ.GetString("SDK_VERSION", "latest"),
		"version of sdk (launchpad, stargate-40, stargate-44 or latest)",
	)
	flag.StringVar(&keyName, "account-name",
		environ.GetString("ACCOUNT_NAME", cosmosfaucet.DefaultAccountName),
		"name of the account to be used by the faucet",
	)
	flag.StringVar(&keyMnemonic, "mnemonic",
		environ.GetString("MNEMONIC", ""),
		"mnemonic for restoring an account",
	)
	flag.StringVar(&keyringPassword, "keyring-password",
		environ.GetString("KEYRING_PASSWORD", ""),
		"password for accessing keyring",
	)
	flag.StringVar(&appCli, "cli-name",
		environ.GetString("CLI_NAME", "nibid"),
		"name of the cli executable",
	)
	flag.StringVar(&defaultDenoms, "denoms",
		environ.GetString("DENOMS", cosmosfaucet.DefaultDenom),
		"denomination of the coins sent by default (comma separated)",
	)
	flag.StringVar(&creditAmount,
		"credit-amount",
		environ.GetString("CREDIT_AMOUNT", strconv.FormatUint(cosmosfaucet.DefaultAmount, 10)),
		"amount to credit in each request per denom, e.g. 1000,1000000",
	)
	flag.StringVar(&maxCredit,
		"max-credit", environ.GetString("MAX_CREDIT", strconv.FormatUint(cosmosfaucet.DefaultMaxAmount, 10)),
		"maximum credit per denom, e.g. 10000,1000000000",
	)
	flag.StringVar(&nodeAddress, "node",
		environ.GetString("NODE", ""),
		"address of tendermint RPC endpoint for this chain",
	)
	flag.BoolVar(&legacySendCmd, "legacy-send",
		environ.GetBool("LEGACY_SEND", false),
		"whether to use legacy send command",
	)
	flag.StringVar(&coinType, "coin-type",
		environ.GetString("COIN_TYPE", "118"),
		"registered coin type number for HD derivation (BIP-0044), defaults from (satoshilabs/SLIP-0044)",
	)
	flag.StringVar(&home, "home",
		environ.GetString("HOME", ""),
		"replaces the default home used by the chain",
	)
}
