package eth

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"coinwallet/wallet/eth/factory"
	"math/big"
	"strings"
)

type EtherTx struct {
	ChainId   int64    `json:"chainId"`
	PrvKeyStr string   `json:"prvKeyStr"`
	Nonce     int64    `json:"nonce"`
	Address   string   `json:"address"`
	Amount    *big.Int `json:"amount"`
	GasLimit  int64    `json:"gasLimit"`
	GasPrice  *big.Int `json:"gasPrice"`
	Data      string   `json:"data"`
}


func ABITransfer(toAddr string, amount *big.Int) ([]byte, error) {
	mAbi, err := abi.JSON(strings.NewReader(factory.TokenABI))
	if err != nil {
		return nil, err
	}
	to := common.HexToAddress(toAddr)
	return mAbi.Pack("transfer", to, amount)
}


func SignEthTx(ethTx EtherTx) (string, error) {
	signer := types.NewEIP155Signer(big.NewInt(ethTx.ChainId))
	tx := types.NewTransaction(
		uint64(ethTx.Nonce),
		common.HexToAddress(ethTx.Address),
		ethTx.Amount, uint64(ethTx.GasLimit),
		ethTx.GasPrice, common.Hex2Bytes(ethTx.Data),
	)
	prvKey, err := crypto.HexToECDSA(ethTx.PrvKeyStr)
	if err != nil {
		return "", err
	}
	h := signer.Hash(tx)
	sig, err := crypto.Sign(h[:], prvKey)
	if err != nil {
		return "", err
	}
	signedTx, err := tx.WithSignature(signer, sig)
	if err != nil {
		return "", err
	}
	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(data), nil
}

func CreateErc20TransferData(toAddress string, amount *big.Int) ([]byte, error) {
	methodId := crypto.Keccak256Hash([]byte("transfer(address,uint256)"))
	var data []byte
	data = append(data, methodId[:4]...)
	paddedAddress := common.LeftPadBytes(common.HexToAddress(toAddress).Bytes(), 32)
	data = append(data, paddedAddress...)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, paddedAmount...)
	return data, nil
}

