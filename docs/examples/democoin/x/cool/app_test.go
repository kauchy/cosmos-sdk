package cool

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/mock"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	priv1  = ed25519.GenPrivKey()
	pubKey = priv1.PubKey()
	addr1  = sdk.AccAddress(pubKey.Address())

	quizMsg1 = MsgQuiz{
		Sender:     addr1,
		CoolAnswer: "icecold",
	}

	quizMsg2 = MsgQuiz{
		Sender:     addr1,
		CoolAnswer: "badvibesonly",
	}

	setTrendMsg1 = MsgSetTrend{
		Sender: addr1,
		Cool:   "icecold",
	}

	setTrendMsg2 = MsgSetTrend{
		Sender: addr1,
		Cool:   "badvibesonly",
	}

	setTrendMsg3 = MsgSetTrend{
		Sender: addr1,
		Cool:   "warmandkind",
	}
)

// initialize the mock application for this module
func getMockApp(t *testing.T) *mock.App {
	mapp := mock.NewApp()

	RegisterCodec(mapp.Cdc)
	keyCool := sdk.NewKVStoreKey("cool")
	bankKeeper := bank.NewBaseKeeper(mapp.AccountKeeper)
	keeper := NewKeeper(keyCool, bankKeeper, DefaultCodespace)
	mapp.Router().AddRoute("cool", NewHandler(keeper))

	mapp.SetInitChainer(getInitChainer(mapp, keeper, "ice-cold"))

	require.NoError(t, mapp.CompleteSetup(keyCool))
	return mapp
}

// overwrite the mock init chainer
func getInitChainer(mapp *mock.App, keeper Keeper, newTrend string) sdk.InitChainer {
	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		mapp.InitChainer(ctx, req)
		keeper.setTrend(ctx, newTrend)

		return abci.ResponseInitChain{}
	}
}

func TestMsgQuiz(t *testing.T) {
	mapp := getMockApp(t)

	// Construct genesis state
	acc1 := &auth.BaseAccount{
		Address: addr1,
		Coins:   nil,
	}
	accs := []auth.Account{acc1}

	// Initialize the chain (nil)
	mock.SetGenesis(mapp, accs)

	// A checkTx context (true)
	ctxCheck := mapp.BaseApp.NewContext(true, abci.Header{})
	res1 := mapp.AccountKeeper.GetAccount(ctxCheck, addr1)
	require.Equal(t, acc1, res1)

	// Set the trend, submit a really cool quiz and check for reward
	mock.SignCheckDeliver(t, mapp.BaseApp, []sdk.Msg{setTrendMsg1}, []uint64{0}, []uint64{0}, true, true, priv1)
	mock.SignCheckDeliver(t, mapp.BaseApp, []sdk.Msg{quizMsg1}, []uint64{0}, []uint64{1}, true, true, priv1)
	mock.CheckBalance(t, mapp, addr1, sdk.Coins{sdk.NewCoin("icecold", sdk.NewInt(69))})
	mock.SignCheckDeliver(t, mapp.BaseApp, []sdk.Msg{quizMsg2}, []uint64{0}, []uint64{2}, false, false, priv1) // result without reward
	mock.CheckBalance(t, mapp, addr1, sdk.Coins{sdk.NewCoin("icecold", sdk.NewInt(69))})
	mock.SignCheckDeliver(t, mapp.BaseApp, []sdk.Msg{quizMsg1}, []uint64{0}, []uint64{3}, true, true, priv1)
	mock.CheckBalance(t, mapp, addr1, sdk.Coins{sdk.NewCoin("icecold", sdk.NewInt(138))})
	mock.SignCheckDeliver(t, mapp.BaseApp, []sdk.Msg{setTrendMsg2}, []uint64{0}, []uint64{4}, true, true, priv1) // reset the trend
	mock.SignCheckDeliver(t, mapp.BaseApp, []sdk.Msg{quizMsg1}, []uint64{0}, []uint64{5}, false, false, priv1)   // the same answer will nolonger do!
	mock.CheckBalance(t, mapp, addr1, sdk.Coins{sdk.NewCoin("icecold", sdk.NewInt(138))})
	mock.SignCheckDeliver(t, mapp.BaseApp, []sdk.Msg{quizMsg2}, []uint64{0}, []uint64{6}, true, true, priv1) // earlier answer now relevant again
	mock.CheckBalance(t, mapp, addr1, sdk.Coins{sdk.NewCoin("badvibesonly", sdk.NewInt(69)), sdk.NewCoin("icecold", sdk.NewInt(138))})
	mock.SignCheckDeliver(t, mapp.BaseApp, []sdk.Msg{setTrendMsg3}, []uint64{0}, []uint64{7}, false, false, priv1) // expect to fail to set the trend to something which is not cool
}
