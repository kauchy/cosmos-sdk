package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/spf13/cobra"
)

func GetCmdQueryMinter(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "minter",
		Short: "Query minter",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryStore([]byte{0x00}, storeName)
			if err != nil {
				return err
			} else if len(res) == 0 {
				return fmt.Errorf("No Minter found")
			}

			var minter mint.Minter
			cdc.MustUnmarshalBinaryLengthPrefixed(res, &minter)

			var output []byte
			if cliCtx.Indent {
				output, err = cdc.MarshalJSONIndent(minter, "", "  ")
			} else {
				output, err = cdc.MarshalJSON(minter)
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

