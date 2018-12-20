package store

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"strings"

	"github.com/cosmos/cosmos-sdk/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	go_amino "github.com/tendermint/go-amino"
)

const (
	flagPath = "path"
	flagData = "data"
)

func StoreCommand(cdc *go_amino.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "store",
		Short: "Query store data by low level",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			node, err := cliCtx.GetNode()
			if err != nil {
				return err
			}

			data:=viper.GetString(flagData)
			var bytes []byte
			if strings.HasPrefix(data,"0x"){
				bytes,_=hex.DecodeString(data)
			}else{
				bytes=[]byte(data)
			}

			result, err := node.ABCIQuery(viper.GetString(flagPath), bytes)
			if err != nil {
				return err
			}

			valueBz := result.Response.GetValue()
			if len(valueBz) == 0 {
				return errors.New("response empty value")
			}

			val, err := tryDecodeValue(cliCtx.Codec, valueBz)
			if err != nil {
				return err
			}

			var bz []byte

			if cliCtx.Indent {
				bz, err = cdc.MarshalJSONIndent(val, "", "  ")
			} else {
				bz, err = cdc.MarshalJSON(val)
			}

			if err != nil {
				return err
			}

			fmt.Println(string(bz))
			return nil
		},
	}

	cmd.Flags().String(flagPath, "", "store query path")
	cmd.MarkFlagRequired(flagPath)
	cmd.Flags().String(flagData, "", "store query data")
	cmd.MarkFlagRequired(flagData)

	cmd.Flags().Bool(client.FlagIndentResponse, false, "print indent result json")

	cmd.Flags().StringP(client.FlagNode, "n", "tcp://localhost:26657", "Node to connect to")
	viper.BindPFlag(client.FlagNode, cmd.Flags().Lookup(client.FlagNode))
	cmd.Flags().Bool(client.FlagTrustNode, false, "Trust connected full node (don't verify proofs for responses)")
	viper.BindPFlag(client.FlagTrustNode, cmd.Flags().Lookup(client.FlagTrustNode))
	cmd.Flags().String(client.FlagChainID, "", "Chain ID of Tendermint node")
	viper.BindPFlag(client.FlagChainID, cmd.Flags().Lookup(client.FlagChainID))

	return cmd
}

func noPaincRegisterInterface(cdc *go_amino.Codec) {
	defer func() {
		if r := recover(); r != nil {
			//nothing
		}
	}()
	cdc.RegisterInterface((*interface{})(nil), nil)
}

type kvPairResult struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func tryDecodeValue(cdc *go_amino.Codec, bz []byte) (interface{}, error) {
	noPaincRegisterInterface(cdc)

	var vInterface interface{}
	err := cdc.UnmarshalBinaryBare(bz, &vInterface)
	if err == nil {
		return vInterface, nil
	}

	var vKVPair []store.KVPair
	err = cdc.UnmarshalBinaryBare(bz, &vKVPair)
	if err == nil {
		var pairResults []kvPairResult
		for _, pair := range vKVPair {
			val, _ := tryDecodeValue(cdc, pair.Value)
			pairResults = append(pairResults, kvPairResult{
				Key:   string(pair.Key),
				Value: val,
			})
		}
		return pairResults, nil
	}

	var vString string
	err = cdc.UnmarshalBinaryBare(bz, &vString)
	if err == nil {
		return vString, nil
	}

	var vInt int64
	err = cdc.UnmarshalBinaryBare(bz, &vInt)
	if err == nil {
		return vInt, nil
	}

	var vBool bool
	err = cdc.UnmarshalBinaryBare(bz, &vBool)
	if err == nil {
		return vBool, nil
	}

	return nil, errors.New("can't decode value")
}
