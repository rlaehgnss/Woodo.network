package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

var b *blockchain
var once sync.Once

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	Blocks []*Block
}

func (b *Block) getHash() {
	Hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", Hash)
}

func getPrevHash() string {
	totalBlocks := len(GetBlockChain().Blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockChain().Blocks[totalBlocks-1].Hash
}

func createBlock(Data string) *Block {
	newBlock := Block{Data, "", getPrevHash()}
	newBlock.getHash()
	return &newBlock
}

func (b *blockchain) AddBlock(Data string) {
	b.Blocks = append(b.Blocks, createBlock(Data))
}

func GetBlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.Blocks
}
