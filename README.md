# go-chonk
*Injects/Extracts chunk data in PNG images.*

## Description

go-chonk will inject a payload into a PNG image, with the chunkType "puNk". Payload can be encrypted by supplying a file containing a 32-byte key. A warning will be displayed when attempting to use an unencrypted payload.

A PNG image that has been injected into using this program can be decrypted by supplying the PNG file, and a file containing the same 32-byte key. 

## Notes

Currently the payload is inserted between the IDAT and END chunks. It will likely cause optimisation failures if run through a program such as "optipng". More research into appropriate chunk data sizes and ciphertext division are necessary to fully accept this as a true stenographic program. It's highly unlikely the data added as part of the current algorithm will draw attention to itself, however, the possibility for it to be affected by optimisation algorithms is considered a destruction vector. Currently, my success criteria are as follows;

1. Store ciphertext in a multi-part fasion within a yet-to-be-determined chunk type
2. Be able to accurately concatenate the ciphertext from it's multi-part form.
3. Ensure data integrity is retained when optimization occurs.
4. Further obfuscate payload with randomised positioning within byte structure.

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
  -target string
        target file
```

