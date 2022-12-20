package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v9/x/reward/types"
)

const ADMIN = "cascadia15e42sv6wm7ds69yy33h55pxr9wjxtzgnf06yxs"

func (k msgServer) RegisterVeContractReward(goCtx context.Context, msg *types.MsgRegisterVeContractReward) (*types.MsgRegisterRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Handling the message
	params := k.GetParams(ctx)
	if !params.EnableReward {
		return nil, types.ErrRewardEnable
	}

	if msg.Creator != ADMIN {
		return nil, types.ErrUnauthorized
	}

	contract := common.HexToAddress(msg.Contract)

	// contract must already be deployed, to avoid spam registrations
	contractAccount := k.evmKeeper.GetAccountWithoutBalance(ctx, contract)

	if contractAccount == nil || !contractAccount.IsContract() {
		return nil, sdkerrors.Wrapf(
			types.ErrRevenueNoContractDeployed,
			"no contract code found at address %s", msg.Contract,
		)
	}

	k.SetReward(ctx, types.Reward{
		Index:             "vecontract",
		Contract:          msg.Contract,
		GasFeeShares:      msg.GasFeeShares,
		BlockRewardShares: msg.BlockRewardShares,
	})

	return &types.MsgRegisterRewardResponse{}, nil
}

func (k msgServer) RegisterNProtocolReward(goCtx context.Context, msg *types.MsgRegisterNProtocolReward) (*types.MsgRegisterRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	params := k.GetParams(ctx)
	if !params.EnableReward {
		return nil, types.ErrRewardEnable
	}

	if msg.Creator != ADMIN {
		return nil, types.ErrUnauthorized
	}

	contract := common.HexToAddress(msg.Contract)

	// contract must already be deployed, to avoid spam registrations
	contractAccount := k.evmKeeper.GetAccountWithoutBalance(ctx, contract)

	if contractAccount == nil || !contractAccount.IsContract() {
		return nil, sdkerrors.Wrapf(
			types.ErrRevenueNoContractDeployed,
			"no contract code found at address %s", msg.Contract,
		)
	}

	k.SetReward(ctx, types.Reward{
		Index:             "nprotocol",
		Contract:          msg.Contract,
		GasFeeShares:      msg.GasFeeShares,
		BlockRewardShares: msg.BlockRewardShares,
	})

	return &types.MsgRegisterRewardResponse{}, nil
}
