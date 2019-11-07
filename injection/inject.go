package injection

import (
	"math/rand"
	"time"
)

func Inject(chunks []Chunk, payload Chunk, shuffle bool) []Chunk {
	var burnedChunks []Chunk

	// randomise position if shuffle, otherwise add after position 2
	pos := 2
	if shuffle {
		rand.Seed(time.Now().Unix())
		pos = rand.Intn(len(chunks))
	}

	for i := range chunks {
		if i == pos {
			burnedChunks = append(burnedChunks, payload)
		}
		burnedChunks = append(burnedChunks, chunks[i])

	}
	return burnedChunks
}
