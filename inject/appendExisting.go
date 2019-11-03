// Appends payload to existing byte structure.
package inject

import (
	"bytes"
	"encoding/binary"
	"os"
)

type Chunk struct {
	Length int    // chunk data length
	CType  []byte // chunk type
	Data   []byte // chunk data
	Crc32  []byte // CRC32 of chunk data
}

func basicInjection() {
	var nothing []byte
	buf := bytes.NewBuffer(nothing)

	var c Chunk

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(c.Length))

	c.Data = []byte("Data to insert")
	c.Length = len(c.Data)
	c.CType = []byte("Datatype")
	c.Crc32 = make([]byte, 4)

	outLength := make([]byte, 4)
	binary.LittleEndian.PutUint32(outLength, uint32(c.Length))

	buf.Write(outLength)
	buf.Write(c.CType)
	buf.Write(c.Data)
	buf.Write(c.Crc32)

	f, err := os.OpenFile("image1.png", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(buf.Bytes())
}
