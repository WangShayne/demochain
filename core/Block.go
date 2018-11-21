package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	// head
	Index        int64  //区块编号
	Timestamp    int64  //时间戳
	PreBlockHash string // 上一个区块链的hash值
	Hash         string // 当前区块的hash值

	// body
	Data string //区块数据
}

// 计算区块hash值
func calculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + string(b.PreBlockHash) + string(b.Data)
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

// 创建区块
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PreBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// 创建创世区块 -- 复用创建区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "hello block")
}
