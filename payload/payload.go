package payload

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"io"
	"log"

	"github.com/DiscoRiver/go-chonk/injection"
)

func BuildPayload(data string, dataType string, shuffle bool) []injection.Chunk {
	var c []injection.Chunk
	var ch injection.Chunk

	if shuffle {
		dataBytes := []byte(data)
		tmp := make([]byte, len(dataBytes)/2)
		tmp2 := make([]byte, len(dataBytes)/2)

		data := bytes.NewReader(dataBytes)

		var chunkBytes [][]byte
		// 1
		if _, err := io.ReadFull(data, tmp); err != nil {
			log.Fatalln(err)
		}
		chunkBytes = append(chunkBytes, tmp)

		// 2
		if _, err := io.ReadFull(data, tmp2); err != nil {
			log.Fatalln(err)
		}
		chunkBytes = append(chunkBytes, tmp2)
		for i := range chunkBytes {
			//index := make([]byte, 5)
			//index = append(index, byte(i))

			//ch.Data = append(index, chunkBytes[i]...)
			ch.Data = chunkBytes[i]
			ch.Length = make([]byte, 4)
			ch.CType = []byte(dataType)
			ch.Crc32 = make([]byte, 4)

			binary.BigEndian.PutUint32(ch.Length, uint32(len(ch.Data)))

			// CRC32
			var crcCheck []byte
			crcCheck = append(crcCheck, append(ch.CType, ch.Data...)...)

			binary.BigEndian.PutUint32(ch.Crc32, crc32.ChecksumIEEE(crcCheck))

			c = append(c, ch)
		}

	} else {

		ch.Data = []byte(data)
		ch.Length = make([]byte, 4)
		ch.CType = []byte(dataType)
		ch.Crc32 = make([]byte, 4)

		binary.BigEndian.PutUint32(ch.Length, uint32(len(ch.Data)))

		// CRC32
		var crcCheck []byte
		crcCheck = append(crcCheck, append(ch.CType, ch.Data...)...)

		binary.BigEndian.PutUint32(ch.Crc32, crc32.ChecksumIEEE(crcCheck))

		c = append(c, ch)
	}

	return c
}
