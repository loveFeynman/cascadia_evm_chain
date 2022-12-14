package keeper

import (
	"github.com/evmos/evmos/v9/x/reward/types"
)

var _ types.QueryServer = Keeper{}
