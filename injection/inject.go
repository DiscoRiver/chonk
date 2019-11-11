package injection

import (
	"math/rand"
	"time"
)

// Inject will build a Chunk slice consisting of the source image chunks, and the payload. Payload position is randomised.
func Inject(chunks []Chunk, payload []Chunk) []Chunk {
	var burnedChunks []Chunk

	// randomise position as standard, otherwise split ciphertext if shuffle
	rand.Seed(time.Now().Unix())
	// restricting the positioning to avoid the header chunk
	pos := rand.Intn(len(chunks)-1) + 1

	burned := false
	for i := range chunks {
		// IDAT structure cannot be broken, must place before or after it.
		if i >= pos && string(chunks[i].CType) != "IDAT" && burned == false {
			for i := range payload {
				burnedChunks = append(burnedChunks, payload[i])
				burned = true
			}
		}
		burnedChunks = append(burnedChunks, chunks[i])
	}
	return burnedChunks
}
