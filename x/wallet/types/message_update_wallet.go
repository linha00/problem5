package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateWallet{}

func NewMsgUpdateWallet(creator string, balance uint64, id uint64) *MsgUpdateWallet {
	return &MsgUpdateWallet{
		Creator: creator,
		Balance: balance,
		Id:      id,
	}
}

func (msg *MsgUpdateWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
