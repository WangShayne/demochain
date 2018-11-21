package core

import (
	"fmt"

	"github.com/cloudflare/cfssl/log"
)

// 区块链结构体
type Blockschain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockschain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockschain{}
	blockchain.AppendNewBlock(&genesisBlock)
	return &blockchain
}

//  发送数据新增区块
func (bc *Blockschain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendNewBlock(&newBlock)
}

// 向区块链结构体中新增一个区块
func (bc *Blockschain) AppendNewBlock(newBlock *Block) {
	// 1.判断区块链中是否已有区块,如果有直接append
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	// 2.调用isValid验证区块是否合法
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

// 打印区块链中的数据
func (bc *Blockschain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("index : %d\n", block.Index)
		fmt.Printf("time : %d\n", block.Timestamp)
		fmt.Printf("preHash : %s\n", block.PreBlockHash)
		fmt.Printf("hash : %s\n", block.Hash)
		fmt.Printf("data : %s\n", block.Data)
		fmt.Println("\n\n")
	}
}

// 验证区块是否为合法的区块
func isValid(newBlock Block, oldBlock Block) bool {
	// 1.判断区块的序列号是否比上一个区块大1
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}
	// 2.判断区块的preHash值是否为上一个区块的hash值
	if newBlock.PreBlockHash != oldBlock.Hash {
		return false
	}
	// 3.判断区块的hash值是否争取
	if newBlock.Hash != calculateHash(newBlock) {
		return false
	}
	return true
}
