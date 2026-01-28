package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/0xSemantic/zonn/x/identity/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) CreateProfile(goCtx context.Context, msg *types.MsgCreateProfile) (*types.MsgCreateProfileResponse, error) {
	return k.Keeper.CreateProfile(goCtx, msg)
}

func (k msgServer) LinkWallet(goCtx context.Context, msg *types.MsgLinkWallet) (*types.MsgLinkWalletResponse, error) {
	return k.Keeper.LinkWallet(goCtx, msg)
}

func (k msgServer) UpdateProfile(goCtx context.Context, msg *types.MsgUpdateProfile) (*types.MsgUpdateProfileResponse, error) {
	return k.Keeper.UpdateProfile(goCtx, msg)
}