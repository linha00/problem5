package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteWallet{}

func NewMsgDeleteWallet(creator string, id uint64) *MsgDeleteWallet {
	return &MsgDeleteWallet{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeleteWallet) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
