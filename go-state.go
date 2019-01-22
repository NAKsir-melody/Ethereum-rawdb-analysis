package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/core/state"
)

func main() {
//state roots
	hash_list := []string{
		"56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
		"1c7677921818add9604af94c6af6cb562adbc95cfdea847f936511e28d8e2897",
		"337e249c268401079fc728c58142710845805285dbc90e7c71bb1b79b9d7a745",
		"0d9348243d7357c491e6a61f4b1305e77dc6acacdb8cc708e662f6a9bab6ca02",
	}

	addr_list := []string{
		"cea8f2236efa20c8fadeb9d66e398a6532cca6c8",
		"8e64566b5eb8f595f7eb2b8d302f2e5613cb8bae",
	}

	ldb, err := ethdb.NewLDBDatabase("chaindata_tx", 0, 0);
	if err != nil {
		fmt.Printf("open error\n" + err.Error())
		return
	}

//state
	stateDB := state.NewDatabase(ldb)
	mystate, err := state.New(common.HexToHash(hash_list[2]), stateDB)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
		return
	}

	addr := common.HexToAddress(addr_list[0])
	if mystate.Exist(addr) {
		balance := mystate.GetBalance(addr)
		nonce := mystate.GetNonce(addr)
		fmt.Printf("%d\n",balance)
		fmt.Printf("%x\n",nonce)
	} else {
		fmt.Printf("no exist\n")
	}
}

