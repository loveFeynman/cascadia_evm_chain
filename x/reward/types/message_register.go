package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ethermint "github.com/evmos/ethermint/types"
)

const TypeMsgRegisterReward = "register"

var _ sdk.Msg = &MsgRegisterVeContractReward{}

func NewMsgRegisterVeContractReward(creator string, contract string, gasFeeShares sdk.Dec, blockRewardShares sdk.Dec) *MsgRegisterVeContractReward {
	return &MsgRegisterVeContractReward{
		Creator:           creator,
		Contract:          contract,
		GasFeeShares:      gasFeeShares,
		BlockRewardShares: blockRewardShares,
	}
}

func (msg *MsgRegisterVeContractReward) Route() string {
	return RouterKey
}

func (msg *MsgRegisterVeContractReward) Type() string {
	return TypeMsgRegisterReward
}

func (msg *MsgRegisterVeContractReward) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterVeContractReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func validateShares(i interface{}) error {
	v, ok := i.(sdk.Dec)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("invalid parameter: nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("value cannot be negative: %T", i)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("value cannot be greater than 1: %T", i)
	}

	return nil
}

func (msg *MsgRegisterVeContractReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := ethermint.ValidateNonZeroAddress(msg.Contract); err != nil {
		return sdkerrors.Wrapf(err, "invalid contract address %s", msg.Contract)
	}

	if err := validateShares(msg.GasFeeShares); err != nil {
		return err
	}

	if err := validateShares(msg.BlockRewardShares); err != nil {
		return err
	}

	return nil
}

func NewMsgRegisterNProtocolReward(creator string, contract string, gasFeeShares sdk.Dec, blockRewardShares sdk.Dec) *MsgRegisterNProtocolReward {
	return &MsgRegisterNProtocolReward{
		Creator:           creator,
		Contract:          contract,
		GasFeeShares:      gasFeeShares,
		BlockRewardShares: blockRewardShares,
	}
}

func (msg *MsgRegisterNProtocolReward) Route() string {
	return RouterKey
}

func (msg *MsgRegisterNProtocolReward) Type() string {
	return TypeMsgRegisterReward
}

func (msg *MsgRegisterNProtocolReward) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterNProtocolReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterNProtocolReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := ethermint.ValidateNonZeroAddress(msg.Contract); err != nil {
		return sdkerrors.Wrapf(err, "invalid contract address %s", msg.Contract)
	}

	if err := validateShares(msg.GasFeeShares); err != nil {
		return err
	}

	if err := validateShares(msg.BlockRewardShares); err != nil {
		return err
	}

	return nil
}
