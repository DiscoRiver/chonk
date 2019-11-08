package extraction

import (
	"github.com/DiscoRiver/go-chonk/injection"
)

func Pull(chunks []injection.Chunk) string {
	var payloadByte []byte

	for i := range chunks {
		if string(chunks[i].CType) == "puNK" {
			// Currently chunks are placed in order, so ciphertext can be rebuilt
			// by using the standard order in which the chunks appear.
			payloadByte = append(payloadByte, chunks[i].Data...)
		}
	}
	return string(payloadByte)
}
