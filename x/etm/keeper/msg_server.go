package keeper

import (
	"context"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	"github.com/rollchains/eventchain/x/etm/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}

	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// CreateEvent implements types.MsgServer.
func (ms msgServer) CreateEvent(ctx context.Context, msg *types.MsgCreateEvent) (*types.MsgCreateEventResponse, error) {
	if err := ms.k.EventMapping.Set(ctx, msg.Organizer, msg.Name); err != nil {
		return nil, err
	}
	return &types.MsgCreateEventResponse{}, nil
}
