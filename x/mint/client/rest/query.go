package rest

import (
	"github.com/cosmos/cosmos-sdk/x/mint"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gorilla/mux"
)

// register REST routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(
		"/mint/minter",
		QueryMinterRequestHandlerFn(storeName, cdc, cliCtx),
	).Methods("GET")
}

func QueryMinterRequestHandlerFn(
	storeName string, cdc *codec.Codec,
	cliCtx context.CLIContext,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := cliCtx.QueryStore([]byte{0x00}, storeName)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var minter mint.Minter
		cdc.MustUnmarshalBinaryLengthPrefixed(res, &minter)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.PostProcessResponse(w, cdc, minter, cliCtx.Indent)
	}
}