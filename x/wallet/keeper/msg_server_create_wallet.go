package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"wallet/x/wallet/types"
)

func (k msgServer) CreateWallet(goCtx context.Context, msg *types.MsgCreateWallet) (*types.MsgCreateWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var wallet = types.Wallet{
		Creator: msg.Creator,
		Balance: msg.Balance,
	}
	id := k.AppendWallet(ctx, wallet)
	return &types.MsgCreateWalletResponse{
		Id: id,
	}, nil
}
