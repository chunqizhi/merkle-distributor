package distributor

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestRBalanceTree(t *testing.T) {
	balances := []RBalance{
		{Account: common.HexToAddress("0xC28969F24Ee259EB9dd1e9008d91fc5cF2def161"), Amount: big.NewInt(1)},
		{Account: common.HexToAddress("0xEee8b9dBda2E4949962B87B816a764Fc4e455E70"), Amount: big.NewInt(2)},
		{Account: common.HexToAddress("0x47D03FC0c0917E2EfF5a8BbB00D3Ad93870AB1be"), Amount: big.NewInt(3)},
		{Account: common.HexToAddress("0xb8c50a9BAE31ce15f2d51f11fa83cFB114B30Ad4"), Amount: big.NewInt(4)},
		{Account: common.HexToAddress("0x3d03297bd4ff9b056900196AAC935C5f58556528"), Amount: big.NewInt(5)},
		{Account: common.HexToAddress("0xb5713Dd2b43cDEEc28dFF4f751Ab24E44996347C"), Amount: big.NewInt(6)},
	}
	tree, err := RNewBalanceTree(balances)
	assert.Nil(t, err)

	// fmt.Println(tree.RGetProof(balances[0].Account, balances[0].Amount))
	// fmt.Println(tree.RGetProof(balances[1].Account, balances[1].Amount))
	// fmt.Println(tree.RGetProof(balances[2].Account, balances[2].Amount))
	// fmt.Println(tree.RGetProof(balances[3].Account, balances[3].Amount))
	// fmt.Println(tree.RGetProof(balances[4].Account, balances[4].Amount))
	// fmt.Println(tree.RGetProof(balances[5].Account, balances[5].Amount))
	// return
	root := tree.RGetRoot()
	assert.Equal(t, root.Hex(), "0x282a769b45c81edf1b08901a662f975fe3adba1710bade83f2c23cc0de38f76e")

	proofs := [][]string{
		{"0x2666a85d2cd3b83406d8de14f04034a169eef61a4c58426b5708cfd338388b98", "0xdfae16f5718d7805be07a99f7003a0c80249da777a8cbfebc0ba3693afe2e769", "0xa9c9b24651c1b8ab4ca5cc918ba7dcab7e84419a96461457c78275796280a394"},
		{"0xe4494a9a20ff17b60c38896eef95a788be99f286255fbbad4bafd228a1f0cbc2", "0xe873ee15fbe7d53d351fcfa7db231ee61773c14218e6430f87badb1920e6cbf4"},
		{"0xb08ab2f12d6dc9fb9cacb5ba1f2590fd4bdb7b6c9bb9503cc207e01081657d3d", "0xe873ee15fbe7d53d351fcfa7db231ee61773c14218e6430f87badb1920e6cbf4"},
		{"0xacd6c9a380439a6f6324f09df08a654ec5b4c9c8c202abe58c473005e5619a1e", "0xf8f7d7c59cfe635268a5b369e8c8c25b3425e05e4f58f115effde47dbfb3e32b", "0xa9c9b24651c1b8ab4ca5cc918ba7dcab7e84419a96461457c78275796280a394"},
		{"0x252518a2d79f0c9f082f6325ae510ecab4bcd3037d22b5fa864a0fe919a254cc", "0xdfae16f5718d7805be07a99f7003a0c80249da777a8cbfebc0ba3693afe2e769", "0xa9c9b24651c1b8ab4ca5cc918ba7dcab7e84419a96461457c78275796280a394"},
		{"0x85c0a0f9b1351c84ab3076753e09baad2c23fe611534c024be4cb356f89ee308", "0xf8f7d7c59cfe635268a5b369e8c8c25b3425e05e4f58f115effde47dbfb3e32b", "0xa9c9b24651c1b8ab4ca5cc918ba7dcab7e84419a96461457c78275796280a394"},
	}
	for idx, balance := range balances {
		p := make(Elements, 0)
		for _, s := range proofs[idx] {
			p = append(p, common.HexToHash(s))
		}
		fmt.Println(balance)
		fmt.Println(p)
		fmt.Println()
		assert.True(t, RVerifyProof(balance.Account, balance.Amount, p, root))
	}

	// info, err := ParseBalanceMap(balances)
	// assert.Nil(t, err)
	// assert.Equal(t, info.MerkleRoot.Hex(), "0x2ec9c2fc2a55df417ba88ecd833f165fa3c5941772ebaf8c5f4debe33f4d1b12")
	// assert.Equal(t, info.TokenTotal, "0x2ee")
}
