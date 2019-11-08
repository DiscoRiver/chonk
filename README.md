# go-chonk
*Injects/Extracts chunk data in PNG images.*

## Description

go-chonk will inject a payload into a PNG image, with the chunkType "puNK". Payload can be encrypted by supplying a file containing a 32-byte key. A warning will be displayed when attempting to use an unencrypted payload.

A PNG image that has been injected into using this program can be decrypted by supplying the PNG file, and a file containing the same 32-byte key. 

## Notes

Current conditions for payload positioning are;

1. After IHDR chunk
2. Before/After IDAT chunks, but not inbetween

Programs such as optipng will be able to run optimisations without causing problems with decryption later. However, some hosting platforms will have their own optimization algorithms and may strip out non-critical chunks from the images for security/storage reasons.

## Example

### Inject
```
$ go-chonk inject -file assets/padlock.png -payload assets/sample-payload-text.txt -target output.png -key assets/keyfile.txt
```

### Extract
```
$ go-chonk extract -file output.png -key assets/keyfile.txt
```
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
 -shuffle
        Shuffle payload position
  -target string
        target file
```

