package shcrypto

import (
	"encoding/binary"
	"math/rand"
)

// computeSeed computes a seed value from the EpochSecretKey.
func computeSeed(key *EpochSecretKey) int64 {
	keyBytes := key.Marshal()
	keyHash := keccak256(keyBytes)
	return int64(binary.LittleEndian.Uint64(keyHash[:8]))
}

// Shuffle shuffles the order of the transactions using the epoch secret key as a source of
// randomness.
func Shuffle(txs [][]byte, key *EpochSecretKey) [][]byte {
	shuffledTxs := append([][]byte{}, txs...)
	swap := func(i, j int) {
		shuffledTxs[i], shuffledTxs[j] = shuffledTxs[j], shuffledTxs[i]
	}
	// XXX Using go's stdlib for shuffling may mean the result of this function changes when
	// using a different go version
	r := rand.New(rand.NewSource(computeSeed(key))) //nolint:gosec
	r.Shuffle(len(shuffledTxs), swap)
	return shuffledTxs
}
