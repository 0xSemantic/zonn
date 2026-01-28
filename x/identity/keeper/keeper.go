package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper as authkeeper" // For dep auth
	"github.com/google/uuid"
	"github.com/0xSemantic/zonn/x/identity/types"
)

type Keeper struct {
	cdc codec.BinaryCodec
	storeKey storetypes.StoreKey
	memKey storetypes.StoreKey
	authKeeper authkeeper.AccountKeeper // For address validation
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	authKeeper authkeeper.AccountKeeper,
) *Keeper {
	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		authKeeper: authKeeper,
	}
}

// CreateProfile creates a new profile.
func (k Keeper) CreateProfile(goCtx context.Context, msg *types.MsgCreateProfile) (*types.MsgCreateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate creator address
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	// Check if profile exists for primary address
	if _, found := k.GetProfileByWallet(ctx, msg.Creator); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "profile already exists for this address")
	}

	// Generate profile ID
	profileID := uuid.New().String()

	profile := types.Profile{
		ProfileId:       profileID,
		PrimaryAddress:  msg.Creator,
		LinkedAddresses: []string{},
		Username:        msg.Username,
		MetadataUri:     msg.MetadataUri,
		CreatedAt:       ctx.BlockTime().Unix(),
		UpdatedAt:       ctx.BlockTime().Unix(),
	}

	// Store profile
	store := ctx.KVStore(k.storeKey)
	profileStore := prefix.NewStore(store, []byte(types.ProfileKeyPrefix))
	b := k.cdc.MustMarshal(&profile)
	profileStore.Set([]byte(profileID), b)

	// Map primary wallet
	walletStore := prefix.NewStore(store, []byte(types.WalletToProfileKeyPrefix))
	walletStore.Set([]byte(msg.Creator), []byte(profileID))

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"ProfileCreated",
			sdk.NewAttribute("profile_id", profileID),
			sdk.NewAttribute("creator", msg.Creator),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgCreateProfileResponse{ProfileId: profileID}, nil
}

// LinkWallet links a wallet to a profile.
func (k Keeper) LinkWallet(goCtx context.Context, msg *types.MsgLinkWallet) (*types.MsgLinkWalletResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate addresses
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}
	if _, err := sdk.AccAddressFromBech32(msg.WalletAddress); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.WalletAddress)
	}

	profile, found := k.GetProfile(ctx, msg.ProfileId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "profile not found")
	}

	// Only owner can link
	if profile.PrimaryAddress != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only profile owner can link wallets")
	}

	// Check if wallet already mapped
	if _, found := k.GetProfileByWallet(ctx, msg.WalletAddress); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "wallet already linked to a profile")
	}

	// Add linked address
	profile.LinkedAddresses = append(profile.LinkedAddresses, msg.WalletAddress)
	profile.UpdatedAt = ctx.BlockTime().Unix()

	// Update profile
	store := ctx.KVStore(k.storeKey)
	profileStore := prefix.NewStore(store, []byte(types.ProfileKeyPrefix))
	b := k.cdc.MustMarshal(&profile)
	profileStore.Set([]byte(msg.ProfileId), b)

	// Map wallet
	walletStore := prefix.NewStore(store, []byte(types.WalletToProfileKeyPrefix))
	walletStore.Set([]byte(msg.WalletAddress), []byte(msg.ProfileId))

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"WalletLinked",
			sdk.NewAttribute("profile_id", msg.ProfileId),
			sdk.NewAttribute("wallet_address", msg.WalletAddress),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgLinkWalletResponse{}, nil
}

// UpdateProfile updates the profile.
func (k Keeper) UpdateProfile(goCtx context.Context, msg *types.MsgUpdateProfile) (*types.MsgUpdateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate creator
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator)
	}

	profile, found := k.GetProfile(ctx, msg.ProfileId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "profile not found")
	}

	if profile.PrimaryAddress != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only profile owner can update")
	}

	if msg.Username != "" {
		profile.Username = msg.Username
	}
	if msg.MetadataUri != "" {
		profile.MetadataUri = msg.MetadataUri
	}
	profile.UpdatedAt = ctx.BlockTime().Unix()

	// Update store
	store := ctx.KVStore(k.storeKey)
	profileStore := prefix.NewStore(store, []byte(types.ProfileKeyPrefix))
	b := k.cdc.MustMarshal(&profile)
	profileStore.Set([]byte(msg.ProfileId), b)

	// Emit event
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"ProfileUpdated",
			sdk.NewAttribute("profile_id", msg.ProfileId),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgUpdateProfileResponse{}, nil
}

// GetProfile retrieves a profile by ID.
func (k Keeper) GetProfile(ctx sdk.Context, profileID string) (val types.Profile, found bool) {
	store := ctx.KVStore(k.storeKey)
	profileStore := prefix.NewStore(store, []byte(types.ProfileKeyPrefix))

	b := profileStore.Get([]byte(profileID))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetProfileByWallet retrieves a profile by wallet address.
func (k Keeper) GetProfileByWallet(ctx sdk.Context, walletAddr string) (val types.Profile, found bool) {
	store := ctx.KVStore(k.storeKey)
	walletStore := prefix.NewStore(store, []byte(types.WalletToProfileKeyPrefix))

	profileIDBytes := walletStore.Get([]byte(walletAddr))
	if profileIDBytes == nil {
		return val, false
	}

	return k.GetProfile(ctx, string(profileIDBytes))
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&params)
	store.Set([]byte("params"), b)
}

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	store := ctx.KVStore(k.storeKey)
	b := store.Get([]byte("params"))
 if b == nil {
		return types.DefaultParams()
	}

 var params types.Params
	k.cdc.MustUnmarshal(b, &params)
	return params
}