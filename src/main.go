package main

import (
	"bytes"
	"container/list"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const ZerosAmount = 3

type Block struct {
	// Static data
	index     uint64
	data      []byte
	timestamp int64
	prevHash  []byte

	// Computed data
	rnd  uint64
	hash []byte
}

var GenesisBlock = Block{
	index:     0,
	data:      []byte{},
	timestamp: 1597266000000000, // 2020-08-13T00:00:00
	prevHash: []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}, // 32 zeros

	// Computed data
	rnd: 1738891,
	hash: []byte{
		0x00, 0x00, 0x00, 0x90, 0xf1, 0x5e, 0x02, 0x1c, 0xe8, 0xd6, 0xe2, 0x0b, 0x22, 0x34, 0x86, 0x5d,
		0x73, 0x0e, 0x99, 0x7b, 0x10, 0xc6, 0xee, 0x33, 0xae, 0x2d, 0xae, 0xa9, 0x7e, 0x50, 0x34, 0x8c,
	},
}

func checkHash(hash [32]byte) bool {
	for i := 0; i < ZerosAmount; i++ {
		if hash[i] != 0 {
			return false
		}
	}

	return true
}

func (block *Block) CalcTempHash(rnd uint64) [32]byte {
	serializedIndex := []byte(strconv.FormatUint(block.index, 10))
	serializedTimestamp := []byte(strconv.FormatInt(block.timestamp, 10))
	serializedRnd := []byte(strconv.FormatUint(rnd, 10))

	serializedBlock := bytes.Join([][]byte{
		serializedIndex,
		block.data,
		serializedTimestamp,
		block.prevHash,
		serializedRnd,
	}, []byte{})

	return sha256.Sum256(serializedBlock)
}

func (block *Block) CalcHash() [32]byte {
	return block.CalcTempHash(block.rnd)
}

func (block *Block) Mine() {
	for {
		rnd := rand.Uint64()
		hash := block.CalcTempHash(rnd)

		if checkHash(hash) {
			block.rnd = rnd
			block.hash = hash[:]

			break
		}
	}
}

func (block *Block) IsValid() bool {
	actualHash := block.CalcHash()

	return bytes.Compare(actualHash[:], block.hash) == 0
}

type BlockChain struct {
	list *list.List
}

func NewBlockChain() *BlockChain {
	chain := BlockChain{
		list: list.New(),
	}

	chain.list.PushBack(&GenesisBlock)

	return &chain
}

func (chain *BlockChain) AddBlock(data []byte) *Block {
	prevBlock := chain.list.Back().Value.(*Block)
	timestamp := time.Now().UnixNano()

	block := Block{
		index:     prevBlock.index + 1,
		data:      data,
		timestamp: timestamp,
		prevHash:  prevBlock.hash,
	}

	block.Mine()
	chain.list.PushBack(&block)

	return &block
}

func (chain *BlockChain) IsValid() bool {
	var prevBlock *Block

	for e := chain.list.Front(); e != nil; e = e.Next() {
		block := e.Value.(*Block)

		if prevBlock != nil {
			if !block.IsValid() || bytes.Compare(block.prevHash, prevBlock.hash) != 0 {
				return false
			}
		} else if block != &GenesisBlock {
			return false
		}

		prevBlock = block
	}

	return true
}

func main() {
	blockchain := NewBlockChain()
	blockchain.AddBlock([]byte{})
	blockchain.AddBlock([]byte{})
	blockchain.AddBlock([]byte{})

	fmt.Println()
	for e := blockchain.list.Front(); e != nil; e = e.Next() {
		block := e.Value.(*Block)

		fmt.Printf("Index: %d\n", block.index)
		fmt.Printf("Rnd: %d\n", block.rnd)
		fmt.Print("PrevHash: ")
		for _, n := range block.prevHash {
			fmt.Printf("%02x", n)
		}
		fmt.Println()
		fmt.Print("Hash: ")
		for _, n := range block.hash {
			fmt.Printf("%02x", n)
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Printf("Is blockchain valid? %t\n", blockchain.IsValid())
}
