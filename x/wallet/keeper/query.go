package keeper

import (
	"wallet/x/wallet/types"
)

var _ types.QueryServer = Keeper{}
