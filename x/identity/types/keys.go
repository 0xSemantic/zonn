package types

const (
	// ModuleName defines the module name
	ModuleName = "identity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_identity"

	// ProfileKeyPrefix is the prefix for profile store
	ProfileKeyPrefix = "Profile/value/"

	// WalletToProfileKeyPrefix is the prefix for wallet to profile mapping
	WalletToProfileKeyPrefix = "WalletToProfile/value/"
)

// ProfileKey returns the store key to retrieve a Profile by ID
func ProfileKey(profileId string) []byte {
	return []byte(ProfileKeyPrefix + profileId)
}

// WalletToProfileKey returns the store key for wallet to profile ID mapping
func WalletToProfileKey(walletAddr string) []byte {
	return []byte(WalletToProfileKeyPrefix + walletAddr)
}