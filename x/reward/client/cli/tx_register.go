package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/reward/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRegisterVeContractReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-vecontract [contract] [gas-fee-shares] [block-reward-shares]",
		Short: "Broadcast message register",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContract := args[0]
			argGasFeeShares := args[1]
			argBlockRewardShares := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			gasFeeShares, err := sdk.NewDecFromStr(argGasFeeShares)
			if err != nil {
				return err
			}

			blockRewardShares, err := sdk.NewDecFromStr(argBlockRewardShares)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterVeContractReward(
				clientCtx.GetFromAddress().String(),
				argContract,
				gasFeeShares,
				blockRewardShares,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRegisterNProtocolReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-nprotocol [contract] [gas-fee-shares] [block-reward-shares]",
		Short: "Broadcast message register",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argContract := args[0]
			argGasFeeShares := args[1]
			argBlockRewardShares := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			gasFeeShares, err := sdk.NewDecFromStr(argGasFeeShares)
			if err != nil {
				return err
			}

			blockRewardShares, err := sdk.NewDecFromStr(argBlockRewardShares)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterNProtocolReward(
				clientCtx.GetFromAddress().String(),
				argContract,
				gasFeeShares,
				blockRewardShares,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
