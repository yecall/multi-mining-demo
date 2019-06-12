package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"math/rand"
	"time"
)

func randomMerkleRoot() [32]byte {

	var root [32]byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 32; i++ {
		root[i] = byte(r.Intn(256))
	}

	return root

}

func diffToTarget(diff uint64) *big.Int {
	diffBig := new(big.Int).SetUint64(diff)
	target := new(big.Int).Div(pow256, diffBig)
	return target
}

func hashToHashBig(hash [32]byte) *big.Int {
	return new(big.Int).SetBytes(hash[:])
}

func hashToHashHex(hash [32]byte) string {
	return hex.EncodeToString(hash[:])
}

func merkleParent(hash1 [32]byte, hash2 [32]byte) [32]byte {

	h := sha256.New()

	h.Write(hash1[:])
	h.Write(hash2[:])
	hash := h.Sum(nil)

	var hashArr [32]byte
	for i := 0; i < 32; i++ {
		hashArr[i] = hash[i]
	}

	return hashArr
}

func bigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}
