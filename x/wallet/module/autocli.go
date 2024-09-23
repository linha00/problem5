package wallet

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "wallet/api/wallet/wallet"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ShowWallet",
					Use:            "show-wallet [id]",
					Short:          "Query show-wallet",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListWallet",
					Use:            "list-wallet",
					Short:          "Query list-wallet",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateWallet",
					Use:            "create-wallet [balance]",
					Short:          "Send a create-wallet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "balance"}},
				},
				{
					RpcMethod:      "UpdateWallet",
					Use:            "update-wallet [balance] [id]",
					Short:          "Send a update-wallet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "balance"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeleteWallet",
					Use:            "delete-wallet [id]",
					Short:          "Send a delete-wallet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
