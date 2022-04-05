package wallet

import "math/big"

const maxBitLen = 255


func equal(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == 0 }

func gt(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == 1 }

func gte(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) >= 0 }

func lt(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == -1 }

func lte(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) <= 0 }

func add(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Add(i, i2) }

func sub(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Sub(i, i2) }

func mul(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Mul(i, i2) }

func div(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Quo(i, i2) }

func mod(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Mod(i, i2) }

func neg(i *big.Int) *big.Int { return new(big.Int).Neg(i) }

func min(i *big.Int, i2 *big.Int) *big.Int {
	if i.Cmp(i2) == 1 {
		return new(big.Int).Set(i2)
	}

	return new(big.Int).Set(i)
}

func max(i *big.Int, i2 *big.Int) *big.Int {
	if i.Cmp(i2) == -1 {
		return new(big.Int).Set(i2)
	}

	return new(big.Int).Set(i)
}

type Int struct {
	i *big.Int
}

func (i Int) BigInt() *big.Int {
	return new(big.Int).Set(i.i)
}

func NewInt(n int64) Int {
	return Int{big.NewInt(n)}
}

func NewIntFromBigInt(i *big.Int) Int {
	if i.BitLen() > maxBitLen {
		panic("NewIntFromBigInt() out of bound")
	}
	return Int{i}
}

func newIntegerFromString(s string) (*big.Int, bool) {
	return new(big.Int).SetString(s, 0)
}

func NewIntFromString(s string) (res Int, ok bool) {
	i, ok := newIntegerFromString(s)
	if !ok {
		return
	}
	if i.BitLen() > maxBitLen {
		ok = false
		return
	}
	return Int{i}, true
}

func NewIntWithDecimal(n int64, dec int) Int {
	if dec < 0 {
		panic("NewIntWithDecimal() decimal is negative")
	}
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(dec)), nil)
	i := new(big.Int)
	i.Mul(big.NewInt(n), exp)
	if i.BitLen() > maxBitLen {
		panic("NewIntWithDecimal() out of bound")
	}
	return Int{i}
}

