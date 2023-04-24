package storage

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	smath "github.com/ava-labs/avalanchego/utils/math"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/crypto"

	"github.com/ava-labs/hypersdk/examples/tokenvm/utils"
)

func TestMethod(
	ctx context.Context,
	db chain.Database,
	pk crypto.PublicKey,
	asset ids.ID,
	amount uint64,
) error {
	dbKey, bal, err := getBalance(ctx, db, pk, asset)
	if err != nil {
		return err
	}
	nbal, err := smath.Add64(bal, amount)
	if err != nil {
		return fmt.Errorf(
			"%w: could not do test method (asset=%s, bal=%d, addr=%v, amount=%d)",
			ErrInvalidBalance,
			asset,
			bal,
			utils.Address(pk),
			amount,
		)
	}
	// return setBalance(ctx, db, dbKey, nbal)
	return db.Insert(ctx, dbKey, binary.BigEndian.AppendUint64(nil, nbal))

}

// func getBalance(
// 	ctx context.Context,
// 	db chain.Database,
// 	pk crypto.PublicKey,
// 	asset ids.ID,
// ) ([]byte, uint64, error) {
// 	k := PrefixBalanceKey(pk, asset)
// 	bal, err := innerGetBalance(db.GetValue(ctx, k))
// 	return k, bal, err
// }

// // [accountPrefix] + [address] + [asset]
// func PrefixBalanceKey(pk crypto.PublicKey, asset ids.ID) (k []byte) {
// 	k = balancePrefixPool.Get().([]byte)
// 	k[0] = balancePrefix
// 	copy(k[1:], pk[:])
// 	copy(k[1+crypto.PublicKeyLen:], asset[:])
// 	return
// }

// func innerGetBalance(
// 	v []byte,
// 	err error,
// ) (uint64, error) {
// 	if errors.Is(err, database.ErrNotFound) {
// 		return 0, nil
// 	}
// 	if err != nil {
// 		return 0, err
// 	}
// 	return binary.BigEndian.Uint64(v), nil
// }

// func setBalance(
// 	ctx context.Context,
// 	db chain.Database,
// 	dbKey []byte,
// 	balance uint64,
// ) error {
// 	return db.Insert(ctx, dbKey, binary.BigEndian.AppendUint64(nil, balance))
// }
