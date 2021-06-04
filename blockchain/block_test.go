package blockchain

import (
	"bytes"
	"math/rand"
	"os"
	"testing"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}

func setup() {
	rand.Seed(time.Now().UnixNano())
}

func shutdown() {
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestCreateBlock(t *testing.T) {
	data := randString(10)
	prevHash := randBytes(12)
	t.Logf("start creating Block with data: '%s' and prevHash: '%v'", data, prevHash)
	block := CreateBlock(data, prevHash)
	if block.Hash == nil {
		t.Errorf("block.Hash must not be nil")
	}
	if !bytes.Equal(block.Data, []byte(data)) {
		t.Errorf("block.Data must be '%s'", []byte(data))
	}
	if !bytes.Equal(block.Hash[:2], []byte{0, 0}) {
		t.Errorf("block.Hash must start with 2 zeros, but is %v", block.Hash)
	}

	t.Logf("Block.Nonce: '%v'", block.Nonce)
	
}

func TestAddBlock(t *testing.T) {
	data := randString(10)
	blockchain := InitBlockChain()
	blockchain.AddBlock(data)
	if len(blockchain.Blocks) != 2 {
		t.Errorf("blockchain must have 2 blocks, but is has %d", len(blockchain.Blocks))
	}
	if string(blockchain.Blocks[0].Data) != "Genesis" {
		t.Errorf("First block must have 'Genesis' data, but it is %s",string(blockchain.Blocks[0].Data))
	}

	if string(blockchain.Blocks[1].Data) != data {
		t.Errorf("Second block data must be  '%s' , but it is %s",data,string(blockchain.Blocks[1].Data))
	}
}


