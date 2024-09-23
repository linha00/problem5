package keeper

import (
    "context"

    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "wallet/x/wallet/types"
)

func (k Keeper) ShowWallet(goCtx context.Context, req *types.QueryShowWalletRequest) (*types.QueryShowWalletResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    ctx := sdk.UnwrapSDKContext(goCtx)
    wallet, found := k.GetWallet(ctx, req.Id)
    if !found {
        return nil, sdkerrors.ErrKeyNotFound
    }

    return &types.QueryShowWalletResponse{Wallet: wallet}, nil
}
