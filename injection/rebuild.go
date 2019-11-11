package injection

import (
	"bytes"
	"log"
	"os"
)

// Rebuild will rebuild a PNG image from a Chunk slice, to the target file with 0644 permissions.
func Rebuild(chunks []Chunk, target string) {

	outFile, err := os.OpenFile(target, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	var b bytes.Buffer

	// Write the PNG header first
	b.Write(header)
	// Then all our chunks (in the order they were read)
	for i := range chunks {
		b.Write(chunks[i].Length)
		b.Write(chunks[i].CType)
		b.Write(chunks[i].Data)
		b.Write(chunks[i].Crc32)
	}

	outFile.Write(b.Bytes())

}
