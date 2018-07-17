package main

import (
        "fmt"
        "log"
	
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
 )

const (
        //mainnet数据库路径
	dbpath = "/media/lzt/Elements/.ethereum/geth/chaindata"
	blocknum = 1000000
)

func main() {
	db, err := ethdb.NewLDBDatabase(dbpath, 128, 1024)
	if err != nil {
		log.Fatal(err)
	}

	hash := rawdb.ReadCanonicalHash(db, blocknum)
	if hash == (common.Hash{}) {
		log.Fatal("block not exist!")
	}

	fmt.Println(hash)

	header := rawdb.ReadHeader(db, hash, blocknum)
	
	h := common.Hash{}
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, header)
	
	fmt.Println(hw.Sum(h[:0]))
}
