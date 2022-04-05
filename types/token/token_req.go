package token

import (
	"coinwallet/types"
	"github.com/pkg/errors"
)

type SoarchTokenReq struct {
	types.PageSizeData
	ChainId   int64   `json:"chain_id"`
	TokeName  string  `json:"toke_name"`
}

func (this SoarchTokenReq) ParamCheck() (int, error) {
	if this.TokeName == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}