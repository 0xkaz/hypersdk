// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package actions

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
	"github.com/ava-labs/hypersdk/consts"
	"github.com/ava-labs/hypersdk/crypto"
	"github.com/ava-labs/hypersdk/examples/tokenvm/auth"
	"github.com/ava-labs/hypersdk/examples/tokenvm/storage"
	"github.com/ava-labs/hypersdk/utils"
	"log"
)

var _ chain.Action = (*TestAction)(nil)

type TestAction struct {
	// To is the recipient of the [Value].
	To crypto.PublicKey `json:"to"`

	// Asset to transfer to [To].
	Asset ids.ID

	// Amount are transferred to [To].
	Value uint64 `json:"value"`
}

func (t *TestAction) StateKeys(rauth chain.Auth, _ ids.ID) [][]byte {
	return [][]byte{
		storage.PrefixBalanceKey(auth.GetActor(rauth), t.Asset),
		storage.PrefixBalanceKey(t.To, t.Asset),
	}
}

func (t *TestAction) Execute(
	ctx context.Context,
	r chain.Rules,
	db chain.Database,
	_ int64,
	rauth chain.Auth,
	_ ids.ID,
	_ bool,
) (*chain.Result, error) {
	actor := auth.GetActor(rauth)
	log.Printf("actor: %v", actor)
	unitsUsed := t.MaxUnits(r) // max units == units
	if t.Value == 0 {
		return &chain.Result{Success: false, Units: unitsUsed, Output: OutputValueZero}, nil
	}
	// if err := storage.SubBalance(ctx, db, actor, t.Asset, t.Value); err != nil {
	// 	return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	// }
	// if err := storage.AddBalance(ctx, db, t.To, t.Asset, t.Value); err != nil {
	// 	return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	// }
	if err := storage.TestMethod(ctx, db, t.To, t.Asset, t.Value); err != nil {
		return &chain.Result{Success: false, Units: unitsUsed, Output: utils.ErrBytes(err)}, nil
	}
	return &chain.Result{Success: true, Units: unitsUsed}, nil
}

func (*TestAction) MaxUnits(chain.Rules) uint64 {
	// We use size as the price of this transaction but we could just as easily
	// use any other calculation.
	return crypto.PublicKeyLen + consts.IDLen + consts.Uint64Len
}

func (t *TestAction) Marshal(p *codec.Packer) {
	p.PackPublicKey(t.To)
	p.PackID(t.Asset)
	p.PackUint64(t.Value)
}

func UnmarshalTestAction(p *codec.Packer, _ *warp.Message) (chain.Action, error) {
	var testAction TestAction
	p.UnpackPublicKey(false, &testAction.To) // can testAction to blackhole
	p.UnpackID(false, &testAction.Asset)     // empty ID is the native asset
	testAction.Value = p.UnpackUint64(true)
	return &testAction, p.Err()
}

func (*TestAction) ValidRange(chain.Rules) (int64, int64) {
	// Returning -1, -1 means that the action is always valid.
	return -1, -1
}
