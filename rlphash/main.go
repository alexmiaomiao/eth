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

//运行结果
//[142 56 180 219 246 177 31 204 59 157 238 132 251 121 134 226 156 160 160 44 236 216 151 124 22 31 247 51 51 41 104 30]
//[142 56 180 219 246 177 31 204 59 157 238 132 251 121 134 226 156 160 160 44 236 216 151 124 22 31 247 51 51 41 104 30]
