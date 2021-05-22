package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

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

const ZerosAmount = 3

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

func main() {
	GenesisBlock.Mine()

	for _, n := range GenesisBlock.hash {
		fmt.Printf("%02x", n)
	}
	fmt.Println()
	fmt.Println(GenesisBlock.rnd)
	fmt.Println(GenesisBlock.IsValid())
}
