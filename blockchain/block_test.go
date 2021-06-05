package blockchain

import (
	"bytes"
	"math/rand"
	"os"
	"testing"
	"time"

	// "github.com/dgraph-io/badger"
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
}

func TestAddBlock(t *testing.T) {
	data := "Alex"
	blockchain := InitBlockChain()
	blockchain.AddBlock(data)
	// blockchain.Database.View(func(txn *badger.Txn) error {
	// 	item, err := txn.Get([]byte("lh"))
	// 	Handle(err)
	// 	lastHash, err := item.Value()
	// 	if bytes.Equal(lastHash, []byte{}) {
	// 		t.Errorf("blockchain lastHash expected to be %v, but is %v", []byte{}, lastHash)
	// 	}
	// 	return err
	// })

}
