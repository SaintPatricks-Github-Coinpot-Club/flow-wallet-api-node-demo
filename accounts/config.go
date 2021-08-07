package accounts

import (
	"github.com/caarlos0/env/v6"
	"github.com/onflow/flow-go-sdk"
)

// Config struct for account service.
type Config struct {
	AdminAccountAddress   string       `env:"ADMIN_ADDRESS,notEmpty"`
	ChainId               flow.ChainID `env:"CHAIN_ID" envDefault:"flow-emulator"`
	AdminProposerKeyCount uint16       `env:"ADMIN_PROPOSER_COUNT" envDefault:"1"`
}

// ParseConfig parses environment variables to a valid Config.
func ParseConfig() (cfg Config) {
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	return
}
