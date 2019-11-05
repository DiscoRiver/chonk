package payload

import (
	"encoding/binary"
	"hash/crc32"

	"github.com/DiscoRiver/go-chonk/injection"
)

func BuildPayload(data string, dataType string) injection.Chunk {

	var c injection.Chunk

	c.Data = []byte(data)
	c.Length = make([]byte, 4)
	c.CType = []byte(dataType)
	c.Crc32 = make([]byte, 4)

	binary.BigEndian.PutUint32(c.Length, uint32(len(c.Data)))

	var crcCheck []byte
	crcCheck = append(crcCheck, c.CType...)
	crcCheck = append(crcCheck, c.Data...)
	binary.BigEndian.PutUint32(c.Crc32, crc32.ChecksumIEEE(crcCheck))
	return c
}
