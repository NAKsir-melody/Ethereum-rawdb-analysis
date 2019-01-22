package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb"
)

// dump all
func main() {
	ldb, err := ethdb.NewLDBDatabase("chaindata", 0, 0);
	if err != nil {
		fmt.Printf("open error\n" + err.Error())
		return
	}
	iter := ldb.NewIterator()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("[%x, %x]", key, value)
		fmt.Printf("\n")
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Printf("iter error\n" + err.Error())
		return
	}
}

