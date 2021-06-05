package blockchain

import (
	"bytes"
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

func TestNewProof(t *testing.T) {
	data := randString(10)
	prevHash := randBytes(12)
	block := &Block{ Data: []byte(data), PrevHash:  prevHash}
	pow := NewProof(block)
	if pow.Block != block {
		t.Errorf("pow.Block must start must be %v, but is %v", block, pow.Block)
	}
	expectedTarget := new(big.Int)
	expectedTarget, ok := expectedTarget.SetString("441711766194596082395824375185729628956870974218904739530401550323154944", 10)
	if !ok {
		t.Errorf("invalid expected target: %v", expectedTarget)
	}
	if pow.Target.Cmp(expectedTarget) != 0 {
		t.Errorf("pow.Target must be '%d', but is '%d'", expectedTarget, pow.Target)
	}
}

func TestInitData(t *testing.T) {
	data := "Alex"
	prevHash := []byte{1, 2, 3}
	block := &Block{ Data: []byte(data), PrevHash:  prevHash}
	pow := NewProof(block)
	nonce := 1
	powData := pow.InitData(nonce)
	expectedPowdata := []byte{1, 2, 3, 65, 108, 101, 120, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 18}

	if !bytes.Equal(powData, expectedPowdata) {
		t.Errorf("data must be '%v', but is '%v'", expectedPowdata, powData)
	}

}

func TestPOWRun(t *testing.T) {
	data := "Alex"
	prevHash := []byte{1, 2, 3}
	block := CreateBlock(data, prevHash)
	pow := NewProof(block)
	nonce, hash := pow.Run()
	if nonce != 33889 {
		t.Errorf("nonce must be '%d', but is '%d'", 33889, nonce)
	}
	expectedHash := []byte{0, 0, 23, 180, 172, 13, 144, 173, 51, 117, 70, 39, 7, 240, 190, 90, 229, 84, 21, 130, 28, 88, 61, 247, 198, 196, 64, 152, 78, 255, 186, 132}
	if !bytes.Equal(hash, expectedHash) {
		t.Errorf("hash must be '%v', but is '%v'", expectedHash, hash)
	}
}

func TestPOWValidate(t *testing.T) {
	data := "Alex"
	prevHash := []byte{1, 2, 3}
	block := &Block{ Data: []byte(data), PrevHash:  prevHash}
	pow := NewProof(block)
	if pow.Validate() {
		t.Errorf("validate must fail while pow was not executed")
	}
	pow.Run()
	if pow.Validate() {
		t.Errorf("validate must fail while block has no nonce")
	}
	block.Nonce = 33889
	pow = NewProof(block)
	pow.Run()
	if !pow.Validate() {
		t.Errorf("validate must be ok")
	}

}
