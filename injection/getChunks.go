package injection

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DiscoRiver/go-chonk/tools"
)

// GetChunks takes the source file [referenceFile] and breaks down chunks into their sub-components, returns a Chunk slice.
func GetChunks(referenceFile *os.File) []Chunk {

	// Read header. This isn't returned with the Chunk slice.
	header = make([]byte, 8)
	if _, err := io.ReadFull(referenceFile, header); err != nil {
		log.Fatalln(err)
	}
	// Make sure it's a PNG image
	if string(header) != PNGHeader {
		fmt.Printf("Wrong PNG header.\nGot %x - Expected %x\n", header, PNGHeader)
		os.Exit(1)
	}

	var chunks []Chunk
	var err error
	// error will occur when reaching EOF
	for err == nil {
		var c Chunk

		length := make([]byte, 4)
		cType := make([]byte, 4)
		crc32 := make([]byte, 4)

		// chunk length
		_, err = io.ReadFull(referenceFile, length)
		var lengthInt int
		lengthInt, err = tools.UInt32ToInt(length)

		// chunk type
		_, err = io.ReadFull(referenceFile, cType)

		// chunk data
		data := make([]byte, lengthInt)
		_, err = io.ReadFull(referenceFile, data)

		// crc32
		_, err = io.ReadFull(referenceFile, crc32)

		// Send chunk to slice
		c.Length = length
		c.CType = cType
		c.Data = data
		c.Crc32 = crc32

		// Dropping the last empty chunk.
		emptyChunk := make([]byte, 4)
		if !bytes.Equal(c.CType, emptyChunk) {
			chunks = append(chunks, c)
		}
	}
	return chunks
}
