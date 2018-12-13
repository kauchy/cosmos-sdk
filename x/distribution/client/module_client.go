package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	distCmds "github.com/cosmos/cosmos-sdk/x/distribution/client/cli"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	distQueryCmd := &cobra.Command{
		Use:   "dist",
		Short: "Distribution commands for the dist module",
	}
	distQueryCmd.AddCommand(client.GetCommands(
		distCmds.GetCmdQueryFeePool(mc.storeKey, mc.cdc),
		distCmds.GetCmdQueryValidatorDistInfo(mc.storeKey, mc.cdc),
		distCmds.GetCmdQueryValidatorDistInfos(mc.storeKey, mc.cdc),
		distCmds.GetCmdQueryDelegationDistInfos(mc.storeKey, mc.cdc),
	)...)

	return distQueryCmd
	return &cobra.Command{Hidden: true}
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	distTxCmd := &cobra.Command{
		Use:   "dist",
		Short: "Distribution transactions subcommands",
	}

	distTxCmd.AddCommand(client.PostCommands(
		distCmds.GetCmdWithdrawRewards(mc.cdc),
		distCmds.GetCmdSetWithdrawAddr(mc.cdc),
	)...)

	return distTxCmd
}
