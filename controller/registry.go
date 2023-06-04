// controller/registry.go

package controller

import (
	"github.com/ava-labs/avalanchego/utils/wrappers"
	// "github.com/ava-labs/hypersdk/chain"
	// "github.com/ava-labs/hypersdk/codec"

	"tokenvm/actions"
	// "tokenvm/auth"
	"tokenvm/consts"
)

// Setup types
func init() {
	// consts.ActionRegistry = codec.NewTypeParser[chain.Action, *warp.Message]()
	// consts.AuthRegistry = codec.NewTypeParser[chain.Auth, *warp.Message]()

	errs := &wrappers.Errs{}
	errs.Add(
		// when registering new actions, ALWAYS make sure to append at the end.
		consts.ActionRegistry.Register(&actions.Transfer{}, actions.UnmarshalTransfer, false),
		consts.ActionRegistry.Register(&actions.CreateAsset{}, actions.UnmarshalCreateAsset, false),
		consts.ActionRegistry.Register(&actions.MintAsset{}, actions.UnmarshalMintAsset, false),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
