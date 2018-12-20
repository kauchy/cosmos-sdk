package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	mintCmds "github.com/cosmos/cosmos-sdk/x/mint/client/cli"
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
	mintQueryCmd := &cobra.Command{
		Use:   "mint",
		Short: "Distribution commands for the dist module",
	}
	mintQueryCmd.AddCommand(client.GetCommands(
		mintCmds.GetCmdQueryMinter(mc.storeKey, mc.cdc),
	)...)

	return mintQueryCmd

}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {

	return &cobra.Command{Hidden: true}
}
