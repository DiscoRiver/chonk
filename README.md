# go-chonk
*Injects/Extracts chunk data in PNG images.*

## Description

go-chonk will inject a payload into a PNG image, with the chunkType "puNk". Payload can be encrypted by supplying a file containing a 32-byte key. A warning will be displayed when attempting to use an unencrypted payload.

A PNG image that has been injected into using this program can be decrypted by supplying the PNG file, and a file containing the same 32-byte key. 

## Usage

### go-chonk extract
```
Usage of extract:
  -file string
        input file
  -key string
        encryption key file
```

### go-chonk inject
```
Usage of inject:
  -c    print chunks
  -chunks
        print chunks
  -file string
        input file
  -key string
        encryption key file
  -payload string
        payload file
  -target string
        target file
```

