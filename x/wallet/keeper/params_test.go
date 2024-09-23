package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "wallet/testutil/keeper"
	"wallet/x/wallet/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.WalletKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
