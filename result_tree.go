package distributor

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func ResultRankToNode(account common.Address, amount *big.Int) common.Hash {
	paddedAddress := common.LeftPadBytes(account.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	var data []byte
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	hash := crypto.Keccak256Hash(data)
	return hash
}

func ResultRankVerifyProof(account common.Address, amount *big.Int, proof Elements, root common.Hash) bool {
	pair := ResultRankToNode(account, amount)
	for _, item := range proof {
		pair = combinedHash(pair, item)
	}

	return pair == root
}

type ResultRank struct {
	Account common.Address
	Amount  *big.Int
}

type ResultRankTree struct {
	tree *MerkleTree
}

func NewResultRankTree(resultRanks []ResultRank) (*ResultRankTree, error) {
	elements := make(Elements, 0, len(resultRanks))
	for _, resultRank := range resultRanks {
		elements = append(elements, ResultRankToNode(resultRank.Account, resultRank.Amount))
	}

	tree, err := NewMerkleTree(elements)
	if err != nil {
		return nil, err
	}

	return &ResultRankTree{tree: tree}, nil
}

func (r *ResultRankTree) GetRoot() common.Hash {
	return r.tree.GetRoot()
}

func (r *ResultRankTree) GetProof(account common.Address, amount *big.Int) ([]common.Hash, error) {
	return r.tree.GetProof(ResultRankToNode(account, amount))
}
