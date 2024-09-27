package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/rollchains/eventchain/x/etm/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: &p}, nil
}

// GetEvent implements types.QueryServer.
func (k Querier) GetEvent(goCtx context.Context, req *types.QueryGetEventRequest) (*types.QueryGetEventResponse, error) {
	v, err := k.Keeper.EventMapping.Get(goCtx, req.Organizer)
	if err != nil {
		return nil, err
	}

	return &types.QueryGetEventResponse{
		Name: v,
	}, nil
}
