package main

import (
	"bytes"
	"math/big"
)

// hash前 24 bit 为0
const targetBits = 24

// ProofOfWork type
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork create a new PoW
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(target)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}