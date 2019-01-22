package main

import (
	"fmt"
	_ "encoding/hex"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/core/state"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/rlp"
	_ "github.com/ethereum/go-ethereum/accounts"
)

func main() {

	hash_list := []string{
		"0d9348243d7357c491e6a61f4b1305e77dc6acacdb8cc708e662f6a9bab6ca02",
	}

	// emptyState: c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470
	// emptyRoot:  56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421 

	ldb, err := ethdb.NewLDBDatabase("chaindata_tx", 0, 0);
	if err != nil {
		fmt.Printf("open error\n" + err.Error())
		return
	}

	trieDB := trie.NewDatabase(ldb)

	//trie
	state_trie, err := trie.NewSecure(common.HexToHash(hash_list[0]), trieDB, 0) //state
	if err != nil {
		fmt.Printf(err.Error() + "\n")
		return
	}

	it := trie.NewIterator(state_trie.NodeIterator(nil))
	for it.Next() {
		fmt.Printf("[%x, %x]\n", it.Key, it.Value)
	}
}

