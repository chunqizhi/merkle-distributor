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
	assert.Equal(t, root.Hex(), "0xb11c1b9db0fe625973559604d567c3e4648f29843d4e9090235da3efaf9e665f")

	proofs := [][]string{
		{"0xfd1b53b98f4be745fea1d50559d89d66acaa965a25e9a2af6f487979a17e2461", "0xbc62f4a806c01aec89da6d4129c9746ff4dba7b097e42982a92ffae896fdcd15"},
		{"0x7b77345eaa9400947d7bd86a5a62344d215f7b8a27ece17960ebc7e72d93435f", "0xbb78ac4afe5aa780eef1c4f5ef79ec9392ec1ca841973e5ca0cf870a1e5c9621", "0x9838f65a7b83b7cbc0d3eabb5929e54071bb46e92e9a433878228a26be791e7f"},
		{"0x34ed1a99f42ef37521e4b51213b416eb310f5f812782a8a7a8125df93bc1b3f9", "0x5dea403160b71e68b98884a2931a6064d33016537a2b9a38e219575c0522e70d", "0x9838f65a7b83b7cbc0d3eabb5929e54071bb46e92e9a433878228a26be791e7f"},
		{"0xe5abd7ebd2b3d831223de17ef2bc575bb383690a53193f5420f37981a2ffbf50", "0xbc62f4a806c01aec89da6d4129c9746ff4dba7b097e42982a92ffae896fdcd15"},
		{"0x5684bd100deca2515d243e82660ca9b7fe92dc1d5fa259eea5f4753be8d0e5c9", "0x5dea403160b71e68b98884a2931a6064d33016537a2b9a38e219575c0522e70d", "0x9838f65a7b83b7cbc0d3eabb5929e54071bb46e92e9a433878228a26be791e7f"},
		{"0xbda27bfd232387f3b0d33f3baa07d850ebc0b07f4e8243b4596c709420fedcb3", "0xbb78ac4afe5aa780eef1c4f5ef79ec9392ec1ca841973e5ca0cf870a1e5c9621", "0x9838f65a7b83b7cbc0d3eabb5929e54071bb46e92e9a433878228a26be791e7f"},
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
