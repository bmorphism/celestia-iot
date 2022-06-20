package types

import (
	"fmt"
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
)

func TestProtoSandwhichAny(t *testing.T) {
	encCfg := makePaymentEncodingConfig()
	cdc := encCfg.Codec
	ir := encCfg.InterfaceRegistry

	rawSand := Sandwhich{Meat: "beef", Bread: "taco"}

	bzSand, err := cdc.MarshalInterface(&rawSand)
	require.NoError(t, err)

	var sand SandwhichI
	err = cdc.UnmarshalInterface(bzSand, &sand)
	require.NoError(t, err)

	any, err := codectypes.NewAnyWithValue(sand)
	require.NoError(t, err)

	resp := SandwhichResponse{Sandwhich: any}

	var unmarshaledSand SandwhichI
	err = ir.UnpackAny(resp.Sandwhich, &unmarshaledSand)
	require.NoError(t, err)

	fmt.Println(unmarshaledSand.SandwhichType())

	realSandwhich, ok := unmarshaledSand.(*Sandwhich)
	require.True(t, ok)

	fmt.Println(realSandwhich.Bread, realSandwhich.Meat)

}
