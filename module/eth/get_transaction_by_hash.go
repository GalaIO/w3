package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func TransactionByHash(hash common.Hash) *TransactionByHashFactory {
	return &TransactionByHashFactory{hash: hash}
}

type TransactionByHashFactory struct {
	// args
	hash common.Hash

	// returns
	result  *types.Transaction
	returns *types.Transaction
}

func (f *TransactionByHashFactory) Returns(tx *types.Transaction) *TransactionByHashFactory {
	f.returns = tx
	return f
}

// CreateRequest implements the core.RequestCreater interface.
func (f *TransactionByHashFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_getTransactionByHash",
		Args:   []interface{}{f.hash},
		Result: &f.result,
	}, nil
}

// HandleResponse implements the core.ResponseHandler interface.
func (f *TransactionByHashFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}
	*f.returns = *f.result
	return nil
}
