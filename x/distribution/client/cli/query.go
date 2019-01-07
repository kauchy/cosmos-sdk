package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/spf13/cobra"
)

func GetCmdQueryFeePool(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feepool",
		Short: "Query feePool",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryStore(distribution.FeePoolKey, storeName)
			if err != nil {
				return err
			} else if len(res) == 0 {
				return fmt.Errorf("No feepool found")
			}

			var feePool types.FeePool
			cdc.MustUnmarshalBinaryLengthPrefixed(res, &feePool)

			var output []byte
			if cliCtx.Indent {
				output, err = cdc.MarshalJSONIndent(feePool, "", "  ")
			} else {
				output, err = cdc.MarshalJSON(feePool)
			}

			if err != nil {
				return err
			}

			fmt.Println(string(output))

			return nil
		},
	}

	return cmd
}

func GetCmdQueryValidatorDistInfos(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vdis",
		Short: "Query all ValidatorDistInfo",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			resKVs, err := cliCtx.QuerySubspace(distribution.ValidatorDistInfoKey, storeName)

			if err != nil {
				return err
			}

			var vdis []types.ValidatorDistInfo
			for _, kv := range resKVs {
				var vdi types.ValidatorDistInfo
				cdc.MustUnmarshalBinaryLengthPrefixed(kv.Value, &vdi)
				vdis = append(vdis, vdi)
			}

			var output []byte
			output, err =  codec.MarshalJSONIndent(cdc, vdis)

			if err != nil {
				return err
			}

			fmt.Println(string(output))

			return nil
		},
	}
	return cmd
}

func GetCmdQueryValidatorDistInfo(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vdi [operator-addr]",
		Short: "Query A ValidatorDistInfo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryStore(distribution.GetValidatorDistInfoKey(addr), storeName)
			if err != nil {
				return err
			} else if len(res) == 0 {
				return fmt.Errorf("No ValidatorDistInfo found")
			}

			var vdi types.ValidatorDistInfo
			cdc.MustUnmarshalBinaryLengthPrefixed(res, &vdi)

			var output []byte
			if cliCtx.Indent {
				output, err = cdc.MarshalJSONIndent(vdi, "", "  ")
			} else {
				output, err = cdc.MarshalJSON(vdi)
			}

			if err != nil {
				return err
			}

			fmt.Println(string(output))

			return nil
		},
	}
	return cmd
}

func GetCmdQueryDelegationDistInfos(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ddis [delegator-addr]",
		Short: "query all DelegationDistInfos from one delegator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			resKVs, err := cliCtx.QuerySubspace(distribution.GetDelegationDistInfosKey(addr), storeName)
			if err != nil {
				return err
			} else if len(resKVs) == 0 {
				return fmt.Errorf("No DelegationDistInfos found")
			}

			var delegationDistInfos []distribution.DelegationDistInfo
			for _, kv := range resKVs {
				var delegationDistInfo distribution.DelegationDistInfo
				cdc.MustUnmarshalBinaryLengthPrefixed(kv.Value, &delegationDistInfo)
				delegationDistInfos = append(delegationDistInfos, delegationDistInfo)
			}

			output, err := codec.MarshalJSONIndent(cdc, delegationDistInfos)
			if err != nil {
				return err
			}

			fmt.Println(string(output))

			return nil
		},
	}
	return cmd
}

func GetCmdWithdrawResult(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdrawResult [operator-addr]",
		Short: "Query a Validator Rewards With Dec",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr, err := sdk.ValAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			params := distribution.NewQueryValidatorParams(addr)

			bz, err := cdc.MarshalJSON(params)
			res, err := cliCtx.QueryWithData("custom/distr/"+distribution.QueryWithdrawResult, bz)
			if err != nil {
				return err
			}

			var coins types.DecCoins
			cdc.UnmarshalJSON(res, &coins)

			var output []byte
			if cliCtx.Indent {
				output, err = cdc.MarshalJSONIndent(coins, "", "  ")
			} else {
				output, err = cdc.MarshalJSON(coins)
			}

			if err != nil {
				return err
			}

			fmt.Println(string(output))

			return nil
		},
	}

	return cmd
}

