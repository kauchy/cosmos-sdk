package rest

import (
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
)

// register REST routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(
		"/dist/feepool",
		QueryFeePoolRequestHandlerFn(storeName, cdc, cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/dist/vdis",
		QueryValidatorDistInfosRequestHandlerFn(storeName, cdc, cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/dist/vdi/{validatorAddr}",
		QueryValidatorDistInfoRequestHandlerFn(storeName, cdc, cliCtx),
	).Methods("GET")
	r.HandleFunc(
		"/dist/ddis/{delegatorAddr}",
		QueryDelegationDistInfosHandlerFn(storeName, cdc, cliCtx),
	).Methods("GET")
}

func QueryFeePoolRequestHandlerFn(
	storeName string, cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := cliCtx.QueryStore(distribution.FeePoolKey, storeName)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var feePool types.FeePool
		cdc.MustUnmarshalBinaryLengthPrefixed(res, &feePool)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.PostProcessResponse(w, cdc, feePool, cliCtx.Indent)
	}
}

func QueryValidatorDistInfosRequestHandlerFn(
	storeName string, cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resKVs, err := cliCtx.QuerySubspace(distribution.ValidatorDistInfoKey, storeName)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var vdis []types.ValidatorDistInfo
		for _, kv := range resKVs {
			var vdi types.ValidatorDistInfo
			cdc.MustUnmarshalBinaryLengthPrefixed(kv.Value, &vdi)
			vdis = append(vdis, vdi)
		}

		utils.PostProcessResponse(w, cdc, vdis, cliCtx.Indent)
	}
}

func QueryValidatorDistInfoRequestHandlerFn(
	storeName string, cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bech32validatorAddr := vars["validatorAddr"]

		validatorAddr, err := sdk.ValAddressFromBech32(bech32validatorAddr)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := cliCtx.QueryStore(distribution.GetValidatorDistInfoKey(validatorAddr), storeName)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var vdi types.ValidatorDistInfo
		cdc.MustUnmarshalBinaryLengthPrefixed(res, &vdi)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.PostProcessResponse(w, cdc, vdi, cliCtx.Indent)
	}
}

func QueryDelegationDistInfosHandlerFn(
	storeName string, cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bech32delegator := vars["delegatorAddr"]

		delegatorAddr, err := sdk.AccAddressFromBech32(bech32delegator)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		resKVs, err := cliCtx.QuerySubspace(distribution.GetDelegationDistInfosKey(delegatorAddr), storeName)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var ddis []distribution.DelegationDistInfo
		for _, kv := range resKVs {
			var ddi distribution.DelegationDistInfo
			cdc.MustUnmarshalBinaryLengthPrefixed(kv.Value, &ddi)
			ddis = append(ddis, ddi)
		}

		utils.PostProcessResponse(w, cdc, ddis, cliCtx.Indent)
	}
}
