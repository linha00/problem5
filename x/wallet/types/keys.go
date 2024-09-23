package types

const (
	// ModuleName defines the module name
	ModuleName = "wallet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_wallet"

	// WalletKey is used to uniquely identify wallets within the system.
	WalletKey = "Wallet/value/"

	// This key will be used to keep track of the ID of the latest wallet added to the store.
	WalletCountKey = "Wallet/count/"
)

var (
	ParamsKey = []byte("p_wallet")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
