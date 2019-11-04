package payload

import (
	"encoding/binary"

	"github.com/DiscoRiver/go-chonk/injection"
)

func BuildPayload(data string, dataType string) injection.Chunk {

	var c injection.Chunk

	c.Data = []byte(data)
	c.Length = make([]byte, 4)
	c.CType = []byte(dataType)
	c.Crc32 = make([]byte, 4)

	binary.LittleEndian.PutUint32(c.Length, uint32(len(c.Data)))

	return c
}
