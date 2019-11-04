package injection

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DiscoRiver/go-chonk/tools"
)

func GetChunks(referenceFile *os.File) []Chunk {

	// Read header
	header = make([]byte, 8)
	if _, err := io.ReadFull(referenceFile, header); err != nil {
		log.Fatalln(err)
	}
	if string(header) != PNGHeader {
		fmt.Printf("Wrong PNG header.\nGot %x - Expected %x\n", header, PNGHeader)
		os.Exit(1)
	}

	var chunks []Chunk
	var err error
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

		// Send chunk to array element
		c.Length = length
		c.CType = cType
		c.Data = data
		c.Crc32 = crc32

		// Dropping the last empty chunk. This won't affect the resulting
		// image, but it will ensure data integrity.
		emptyChunk := make([]byte, 4)
		if !bytes.Equal(c.CType, emptyChunk) {
			chunks = append(chunks, c)
		}
	}
	return chunks
}
