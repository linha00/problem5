package keeper

import (
    "context"

    "cosmossdk.io/store/prefix"
    "github.com/cosmos/cosmos-sdk/runtime"
    "github.com/cosmos/cosmos-sdk/types/query"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "wallet/x/wallet/types"
)

func (k Keeper) ListWallet(ctx context.Context, req *types.QueryListWalletRequest) (*types.QueryListWalletResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

    storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
    store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WalletKey))

    var wallets []types.Wallet
    pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
        var wallet types.Wallet
        if err := k.cdc.Unmarshal(value, &wallet); err != nil {
            return err
        }

        // Filter wallets based on balance condition
        if wallet.Balance > 100 { // Change 100 to the desired amount
            wallets = append(wallets, wallet)
        }
        return nil
    })

    if err != nil {
        return nil, status.Error(codes.Internal, err.Error())
    }

    return &types.QueryListWalletResponse{Wallet: wallets, Pagination: pageRes}, nil
}
