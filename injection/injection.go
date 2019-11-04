package injection

import "fmt"

const PNGHeader = "\x89\x50\x4E\x47\x0D\x0A\x1A\x0A"

var header []byte

type Chunk struct {
	Length []byte // chunk data length
	CType  []byte // chunk type
	Data   []byte // chunk data
	Crc32  []byte // CRC32 of chunk data
}

// ParseChunk prints chunk data
func PrintChunks(chunks []Chunk) {
	var output string
	for i := range chunks {

		output += fmt.Sprintf("----------\n")
		output += fmt.Sprintf("Chunk length: %d\n", chunks[i].Length)
		output += fmt.Sprintf("Chunk type: %v (%v)\n", chunks[i].CType, string(chunks[i].CType))

		cap := 10
		if len(chunks[i].Data) < 10 {
			cap = len(chunks[i].Data)
		}
		output += fmt.Sprintf("Chunk data (10 bytes): %x\n", chunks[i].Data[:cap])
		output += fmt.Sprintf("----------\n")

	}
	fmt.Print(output)
}
