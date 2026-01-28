package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/0xSemantic/zonn/x/identity/types"
)

func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, profile := range genState.Profiles {
		store := ctx.KVStore(k.storeKey)
		profileStore := prefix.NewStore(store, []byte(types.ProfileKeyPrefix))
		b := k.cdc.MustMarshal(&profile)
		profileStore.Set([]byte(profile.ProfileId), b)

		walletStore := prefix.NewStore(store, []byte(types.WalletToProfileKeyPrefix))
		walletStore.Set([]byte(profile.PrimaryAddress), []byte(profile.ProfileId))

		for _, linked := range profile.LinkedAddresses {
			walletStore.Set([]byte(linked), []byte(profile.ProfileId))
		}
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	profiles := []types.Profile{}
	store := ctx.KVStore(k.storeKey)
	profileStore := prefix.NewStore(store, []byte(types.ProfileKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(profileStore, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var profile types.Profile
		k.cdc.MustUnmarshal(iterator.Value(), &profile)
		profiles = append(profiles, profile)
	}

	return &types.GenesisState{
		Profiles: profiles,
		Params:   params,
	}
}