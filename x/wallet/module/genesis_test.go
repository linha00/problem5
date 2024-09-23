package wallet_test

import (
	"testing"

	keepertest "wallet/testutil/keeper"
	"wallet/testutil/nullify"
	wallet "wallet/x/wallet/module"
	"wallet/x/wallet/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WalletKeeper(t)
	wallet.InitGenesis(ctx, k, genesisState)
	got := wallet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
