package relayer

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/cicizeo/loran/orchestrator/cosmos"
	"github.com/cicizeo/loran/orchestrator/ethereum/peggy"
	"github.com/cicizeo/loran/orchestrator/ethereum/provider"
	"github.com/cicizeo/hilo/x/peggy/types"
)

type PeggyRelayer interface {
	Start(ctx context.Context) error

	FindLatestValset(ctx context.Context) (*types.Valset, error)
	RelayBatches(ctx context.Context) error
	RelayValsets(ctx context.Context) error
}

type peggyRelayer struct {
	logger             zerolog.Logger
	cosmosQueryClient  cosmos.PeggyQueryClient
	peggyContract      peggy.Contract
	ethProvider        provider.EVMProvider
	valsetRelayEnabled bool
	batchRelayEnabled  bool
	loopDuration       time.Duration
}

func NewPeggyRelayer(
	logger zerolog.Logger,
	cosmosQueryClient cosmos.PeggyQueryClient,
	peggyContract peggy.Contract,
	valsetRelayEnabled bool,
	batchRelayEnabled bool,
	loopDuration time.Duration,
) PeggyRelayer {
	return &peggyRelayer{
		logger:             logger.With().Str("module", "peggy_relayer").Logger(),
		cosmosQueryClient:  cosmosQueryClient,
		peggyContract:      peggyContract,
		ethProvider:        peggyContract.Provider(),
		valsetRelayEnabled: valsetRelayEnabled,
		batchRelayEnabled:  batchRelayEnabled,
		loopDuration:       loopDuration,
	}
}
