package distribution

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keep "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
)

// query endpoints supported by the staking Querier
const (
	QuerierRoute 					   = "distr"

	QueryWithdrawResult				   = "withdrawResult"
)

// creates a querier for staking REST endpoints
func NewQuerier(k keep.Keeper, cdc *codec.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryWithdrawResult:
			return queryWithdrawResult(ctx, cdc,req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown dist query endpoint")
		}
	}
}

type QueryValidatorParams struct {
	ValidatorAddr sdk.ValAddress
}

// creates a new QueryValidatorParams
func NewQueryValidatorParams(validatorAddr sdk.ValAddress) QueryValidatorParams {
	return QueryValidatorParams{
		ValidatorAddr: validatorAddr,
	}
}

func queryWithdrawResult(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery, k keep.Keeper) (res []byte, err sdk.Error) {
	var params QueryValidatorParams

	errRes := cdc.UnmarshalJSON(req.Data, &params)
	if errRes != nil {
		return []byte{}, sdk.ErrUnknownAddress("")
	}

	withdraw, _ := k.CurrentValidatorRewardsAllWithDec(ctx, params.ValidatorAddr)

	res, errRes = codec.MarshalJSONIndent(cdc, withdraw)
	if errRes != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", errRes.Error()))
	}
	return res, nil
}
