package types

import (
	"errors"
	"math/big"
	"strings"
)

type BigInt struct {
	Value *big.Int
}

func (b *BigInt) UnmarshalJSON(data []byte) error {
	val := new(big.Int)
	bigNum, ok := val.SetString(strings.Trim(string(data), "\""), 10)
	if !ok {
		return errors.New("cannot unmarshal the field")
	}

	b.Value = bigNum

	return nil
}

func (b *BigInt) String() string {
	return b.Value.String()
}

type Coin struct {
	CoinType        string  `json:"coinType"`
	CoinObjectCount uint64  `json:"coinObjectCount"`
	TotalBalance    *BigInt `json:"totalBalance"`
	// LockedBalance   *big.Int `json:"lockedBalance"` // TODO
}
