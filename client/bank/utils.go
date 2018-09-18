package bank

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// build the sendTx msg
func BuildMsg(from sdk.AccAddress, to sdk.AccAddress, coins sdk.Coins) sdk.Msg {
	input := bank.NewInput(from, coins)
	output := bank.NewOutput(to, coins)
	msg := bank.NewMsgSend([]bank.Input{input}, []bank.Output{output})
	return msg
}