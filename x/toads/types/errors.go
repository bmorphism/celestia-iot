package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrToadIsBigoted = sdkerrors.Register(ModuleName, 11110, "your toad needs to learn some manners")
)
