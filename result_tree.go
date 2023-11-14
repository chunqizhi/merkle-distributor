package distributor

import (
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
	// solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func RToNode(account common.Address, amount *big.Int) common.Hash {
	paddedAddress := common.LeftPadBytes(account.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	var data []byte
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	hash := crypto.Keccak256Hash(data)
	return hash
}

func RVerifyProof(account common.Address, amount *big.Int, proof Elements, root common.Hash) bool {
	pair := RToNode(account, amount)
	for _, item := range proof {
		pair = combinedHash(pair, item)
	}

	return pair == root
}

type RBalance struct {
	Account common.Address
	Amount  *big.Int
}

type RBalanceTree struct {
	rtree *MerkleTree
}

func RNewBalanceTree(balances []RBalance) (*RBalanceTree, error) {
	elements := make(Elements, 0, len(balances))
	for _, balance := range balances {
		elements = append(elements, RToNode(balance.Account, balance.Amount))
	}

	rtree, err := NewMerkleTree(elements)
	if err != nil {
		return nil, err
	}

	return &RBalanceTree{rtree: rtree}, nil
}

func (b *RBalanceTree) RGetRoot() common.Hash {
	return b.rtree.GetRoot()
}

func (b *RBalanceTree) RGetProof(account common.Address, amount *big.Int) ([]common.Hash, error) {
	return b.rtree.GetProof(RToNode(account, amount))
}
