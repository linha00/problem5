package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"wallet/x/wallet/types"
)

func (k Keeper) AppendWallet(ctx sdk.Context, wallet types.Wallet) uint64 {
	count := k.GetWalletCount(ctx)
	wallet.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WalletKey))
	appendedValue := k.cdc.MustMarshal(&wallet)
	store.Set(GetWalletIDBytes(wallet.Id), appendedValue)
	k.SetWalletCount(ctx, count+1)
	return count
}

func (k Keeper) GetWalletCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.WalletCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetWalletIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetWalletCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.WalletCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetWallet(ctx sdk.Context, id uint64) (val types.Wallet, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WalletKey))
	b := store.Get(GetWalletIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetWallet(ctx sdk.Context, wallet types.Wallet) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WalletKey))
	b := k.cdc.MustMarshal(&wallet)
	store.Set(GetWalletIDBytes(wallet.Id), b)
}

func (k Keeper) RemoveWallet(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WalletKey))
	store.Delete(GetWalletIDBytes(id))
}
