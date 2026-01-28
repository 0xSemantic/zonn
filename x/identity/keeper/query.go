package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/0xSemantic/zonn/x/identity/types"
)

// Profile returns a profile by its ID.
func (k Keeper) Profile(goCtx context.Context, req *types.QueryProfileRequest) (*types.QueryProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, found := k.GetProfile(ctx, req.ProfileId)
	if !found {
		return nil, status.Error(codes.NotFound, "profile not found")
	}

	return &types.QueryProfileResponse{Profile: &profile}, nil
}

// ProfileByWallet returns a profile by any associated wallet address.
func (k Keeper) ProfileByWallet(goCtx context.Context, req *types.QueryProfileByWalletRequest) (*types.QueryProfileByWalletResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, found := k.GetProfileByWallet(ctx, req.WalletAddress)
	if !found {
		return nil, status.Error(codes.NotFound, "profile not found")
	}

	return &types.QueryProfileByWalletResponse{Profile: &profile}, nil
}

// Profiles returns all profiles.
func (k Keeper) Profiles(goCtx context.Context, req *types.QueryProfilesRequest) (*types.QueryProfilesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var profiles []types.Profile
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ProfileKeyPrefix))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var profile types.Profile
		k.cdc.MustUnmarshal(iterator.Value(), &profile)
		profiles = append(profiles, profile)
	}

	return &types.QueryProfilesResponse{Profiles: profiles}, nil
}