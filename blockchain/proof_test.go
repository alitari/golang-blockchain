package blockchain

import (
	"math/big"
	"testing"
)

// Take the data from the block

// create a counter (nonce) which starts at 0

// create a hash of the data plus the counter

// check the hash to see if it meets a set of requirements

// Requirements:
// The First few bytes must contain 0s

// const Difficulty = 18

// type ProofOfWork struct {
// 	Block  *Block
// 	Target *big.Int
// }


func TestNewProof(t *testing.T)  {
	data := randString(10)
	prevHash := randBytes(12)
	t.Logf("start creating Block with data: '%s' and prevHash: '%v'", data, prevHash)
	block := CreateBlock(data, prevHash)
	pow := NewProof(block)
	if pow.Block != block {
		t.Errorf("pow.Block must start must be %v, but is %v", block, pow.Block)
	}
	expectedTarget :=  new(big.Int)
	expectedTarget, ok := expectedTarget.SetString("441711766194596082395824375185729628956870974218904739530401550323154944",10)
	if !ok {
		t.Errorf("invalid expected target: %v",expectedTarget)
	}
	if pow.Target.Cmp(expectedTarget) != 0 {
		t.Errorf("pow.Target must be '%d', but is '%d'", expectedTarget ,pow.Target)
	}
}

func TestInitData(t *testing.T)  {
	// data := bytes.Join(
	// 	[][]byte{
	// 		pow.Block.PrevHash,
	// 		pow.Block.Data,
	// 		ToHex(int64(nonce)),
	// 		ToHex(int64(Difficulty)),
	// 	},
	// 	[]byte{},
	// )

	// return data
}

func TestPOWRun(t *testing.T)  {
	
}

func TestPOWValidate(t *testing.T)  {
	
}
