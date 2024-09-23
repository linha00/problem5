package wallet

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"wallet/testutil/sample"
	walletsimulation "wallet/x/wallet/simulation"
	"wallet/x/wallet/types"
)

// avoid unused import issue
var (
	_ = walletsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateWallet = "op_weight_msg_create_wallet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateWallet int = 100

	opWeightMsgUpdateWallet = "op_weight_msg_update_wallet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateWallet int = 100

	opWeightMsgDeleteWallet = "op_weight_msg_delete_wallet"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteWallet int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	walletGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&walletGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateWallet int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateWallet, &weightMsgCreateWallet, nil,
		func(_ *rand.Rand) {
			weightMsgCreateWallet = defaultWeightMsgCreateWallet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateWallet,
		walletsimulation.SimulateMsgCreateWallet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateWallet int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateWallet, &weightMsgUpdateWallet, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateWallet = defaultWeightMsgUpdateWallet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateWallet,
		walletsimulation.SimulateMsgUpdateWallet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteWallet int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteWallet, &weightMsgDeleteWallet, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteWallet = defaultWeightMsgDeleteWallet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteWallet,
		walletsimulation.SimulateMsgDeleteWallet(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateWallet,
			defaultWeightMsgCreateWallet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				walletsimulation.SimulateMsgCreateWallet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateWallet,
			defaultWeightMsgUpdateWallet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				walletsimulation.SimulateMsgUpdateWallet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteWallet,
			defaultWeightMsgDeleteWallet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				walletsimulation.SimulateMsgDeleteWallet(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
