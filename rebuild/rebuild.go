// Rebuilds a .png file at byte level, retaining full byte integrity
package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var filename string

type Chunk struct {
	Length []byte // chunk data length
	CType  []byte // chunk type
	Data   []byte // chunk data
	Crc32  []byte // CRC32 of chunk data
}

func uInt32ToInt(buf []byte) (int, error) {
	if len(buf) == 0 || len(buf) > 4 {
		return 0, errors.New("invalid buffer")
	}
	return int(binary.BigEndian.Uint32(buf)), nil
}

// ParseChunk prints chunk data
func (chunk Chunk) ParseChunk() string {
	var output string

	output += fmt.Sprintf("----------\n")
	output += fmt.Sprintf("Chunk length: %d\n", chunk.Length)
	output += fmt.Sprintf("Chunk type: %v (%v)\n", chunk.CType, string(chunk.CType))

	cap := 10
	if len(chunk.Data) < 10 {
		cap = len(chunk.Data)
	}
	output += fmt.Sprintf("Chunk data (10 bytes): %x\n", chunk.Data[:cap])
	output += fmt.Sprintf("----------\n")

	return output
}

func calcMD5(f *os.File) []byte {
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatalln(err)
	}

	return h.Sum(nil)
}

func init() {
	flag.StringVar(&filename, "file", "", "input file")

	flag.Parse()
}

func main() {
	// File to rebuild
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Read header
	var PNGHeader = "\x89\x50\x4E\x47\x0D\x0A\x1A\x0A"
	header := make([]byte, 8)
	if _, err := io.ReadFull(file, header); err != nil {
		log.Fatalln(err)
	}
	if string(header) != PNGHeader {
		fmt.Printf("Wrong PNG header.\nGot %x - Expected %x\n", header, PNGHeader)
		return
	}

	var chunks []Chunk
	err = nil
	for err == nil {
		var c Chunk

		length := make([]byte, 4)
		cType := make([]byte, 4)
		crc32 := make([]byte, 4)

		// chunk length
		_, err = io.ReadFull(file, length)
		var lengthInt int
		lengthInt, err = uInt32ToInt(length)

		// chunk type
		_, err = io.ReadFull(file, cType)

		// chunk data
		data := make([]byte, lengthInt)
		_, err = io.ReadFull(file, data)

		// crc32
		_, err = io.ReadFull(file, crc32)

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
			fmt.Printf("\nDone processing chunk\n %v", c.ParseChunk())
		}
	}

	// We have our chunks, now just to rebuild the file.
	fmt.Printf("----------\n")
	fmt.Printf("All chunks processed, rebuilding...")

	outFile, err := os.OpenFile("rebuilt.png", os.O_CREATE|os.O_APPEND, 0644)
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

	fmt.Printf("success!\n")
	fmt.Printf("----------\n")

	// Verify the integrity
	fmt.Printf("----------\n")
	fmt.Printf("Checking integrity...")

	if bytes.Equal(calcMD5(file), calcMD5(outFile)) {
		fmt.Printf("verified!\n")
	} else {
		fmt.Printf("failed!\n")
	}
	fmt.Printf("----------\n")

}
