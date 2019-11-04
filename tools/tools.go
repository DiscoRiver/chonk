package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func UInt32ToInt(buf []byte) (int, error) {
	if len(buf) == 0 || len(buf) > 4 {
		return 0, errors.New("invalid buffer")
	}
	return int(binary.BigEndian.Uint32(buf)), nil
}

func CalcMD5(f *os.File) []byte {
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatalln(err)
	}

	return h.Sum(nil)
}

func VerifyIntegrity(f1, f2 *os.File) {
	fmt.Printf("----------\n")
	fmt.Printf("Checking integrity...")

	if bytes.Equal(CalcMD5(f1), CalcMD5(f2)) {
		fmt.Printf("verified!\n")
	} else {
		fmt.Printf("failed!\n")
	}
	fmt.Printf("----------\n")
}
