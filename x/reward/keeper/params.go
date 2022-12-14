package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/reward/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.EnableReward(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// EnableReward returns the EnableReward param
func (k Keeper) EnableReward(ctx sdk.Context) (res bool) {
	k.paramstore.Get(ctx, types.KeyEnableReward, &res)
	return
}
