package payload

import "github.com/DiscoRiver/go-chonk/injection"

func Pull(chunks []injection.Chunk) string {
	var payloadString string
	for i := range chunks {
		if string(chunks[i].CType) == "puNK" {
			payloadString = string(chunks[i].Data)
		}
	}
	return payloadString
}
