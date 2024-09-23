package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"wallet/x/wallet/types"
)

func (k msgServer) UpdateWallet(goCtx context.Context, msg *types.MsgUpdateWallet) (*types.MsgUpdateWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var wallet = types.Wallet{
		Creator: msg.Creator,
		Id:      msg.Id,
		Balance: msg.Balance,
	}
	val, found := k.GetWallet(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetWallet(ctx, wallet)
	return &types.MsgUpdateWalletResponse{}, nil
}
