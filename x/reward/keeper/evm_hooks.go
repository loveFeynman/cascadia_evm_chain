package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var _ evmtypes.EvmHooks = Hooks{}

// Hooks wrapper struct for fees keeper
type Hooks struct {
	k Keeper
}

// Hooks return the wrapper hooks struct for the Keeper
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// PostTxProcessing is a wrapper for calling the EVM PostTxProcessing hook on
// the module keeper
func (h Hooks) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	return h.k.PostTxProcessing(ctx, msg, receipt)
}

// PostTxProcessing implements EvmHooks.PostTxProcessing. After each successful
// interaction with a registered contract, the contract deployer (or, if set,
// the withdraw address) receives a share from the transaction fees paid by the
// transaction sender.
func (k Keeper) PostTxProcessing(
	ctx sdk.Context,
	msg core.Message,
	receipt *ethtypes.Receipt,
) error {
	// check if the fees are globally enabled
	params := k.GetParams(ctx)

	if !params.EnableReward {
		return nil
	}

	txFee := sdk.NewIntFromUint64(receipt.GasUsed).Mul(sdk.NewIntFromBigInt(msg.GasPrice()))
	evmDenom := k.evmKeeper.GetParams(ctx).EvmDenom

	// distribute the fees to the VE contract
	veContract, found := k.GetReward(ctx, "vecontract")
	if found {
		veContractDist := txFee.ToDec().Mul(veContract.GasFeeShares).TruncateInt()
		veContractFees := sdk.Coins{{Denom: evmDenom, Amount: veContractDist}}

		veContractAddress, err := sdk.AccAddressFromHex(veContract.Contract[2:])
		if err != nil {
			goto nProtocol
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			k.feeCollectorName,
			veContractAddress,
			veContractFees,
		)

		if err != nil {
			goto nProtocol
		}
	}

	// distribute the fees to the nProtocol
nProtocol:
	nProtocol, found := k.GetReward(ctx, "nprotocol")

	if !found {
		return nil
	}

	nProtocolDist := txFee.ToDec().Mul(nProtocol.GasFeeShares).TruncateInt()
	nProtocolFees := sdk.Coins{{Denom: evmDenom, Amount: nProtocolDist}}

	nProtocolAddress, err := sdk.AccAddressFromHex(nProtocol.Contract[2:])
	if err != nil {
		return nil
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		k.feeCollectorName,
		nProtocolAddress,
		nProtocolFees,
	)

	if err != nil {
		return sdkerrors.Wrapf(
			err,
			"fee collector account failed to distribute developer fees (%s) to nProtocol address %s.",
			nProtocolFees, nProtocol.Contract,
		)
	}

	return nil
}
