package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/DiscoRiver/go-chonk/extraction"
	"github.com/DiscoRiver/go-chonk/payload"

	"github.com/DiscoRiver/go-chonk/injection"
)

var filename string
var target string
var payloadFile string
var keyFile string
var chunkFlag bool
var shuffle bool

func main() {

	flag.StringVar(&filename, "file", "", "input file")
	flag.BoolVar(&chunkFlag, "c", false, "print chunks")

	ex := flag.NewFlagSet("extract", flag.ContinueOnError)
	ex.StringVar(&filename, "file", "", "input file")
	ex.StringVar(&keyFile, "key", "", "encryption key file")

	in := flag.NewFlagSet("inject", flag.ContinueOnError)
	in.StringVar(&filename, "file", "", "input file")
	in.StringVar(&target, "target", "", "target file")
	in.StringVar(&payloadFile, "payload", "", "payload file")
	in.StringVar(&keyFile, "key", "", "encryption key file")
	in.BoolVar(&chunkFlag, "c", false, "print chunks")
	in.BoolVar(&shuffle, "shuffle", false, "Shuffle payload position")

	switch os.Args[1] {
	case "extract":
		ex.Parse(os.Args[2:])
		if os.Args[2] != "--help" {
			extract()
		}
	case "inject":
		in.Parse(os.Args[2:])
		if os.Args[2] != "--help" {
			inject()
		}
	default:
		flag.Parse()
		if chunkFlag {
			FileStat()
		}
	}
}

func extract() {
	// reference file
	referenceFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer referenceFile.Close()

	referenceChunks := injection.GetChunks(referenceFile)

	payloadChunkString := extraction.Pull(referenceChunks)
	if payloadChunkString == "" {
		fmt.Printf("Payload chunk not found in this file.")
		os.Exit(0)
	}

	var key []byte
	key, err = ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatalln(err)
	}
	plainString := payload.DecryptAES(key, payloadChunkString)

	fmt.Printf("\n----PAYLOAD DECRYPTED----\n")
	fmt.Printf("%v\n", plainString)
	fmt.Printf("----END PAYLOAD----\n")
}

func FileStat() {
	// reference file
	referenceFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer referenceFile.Close()

	// Get reference file chunks
	referenceChunks := injection.GetChunks(referenceFile)
	injection.PrintChunks(referenceChunks)
}

func inject() {
	// reference file
	referenceFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer referenceFile.Close()

	// Get reference file chunks
	referenceChunks := injection.GetChunks(referenceFile)

	if chunkFlag {
		fmt.Printf("----BEGIN REFERNCE FILE CHUNKS----\n")
		injection.PrintChunks(referenceChunks)
		fmt.Printf("----END REFERNCE FILE CHUNKS----\n\n")
	}

	// Get the plaintext payload
	var payloadString string
	if payloadFile != "" {
		var payloadByte []byte
		payloadByte, err = ioutil.ReadFile(payloadFile)
		if err != nil {
			log.Fatalln(err)
		}
		payloadString = string(payloadByte)
	} else {
		fmt.Println("No payload file given, nothing to do.")
		os.Exit(0)
	}

	// Get the key
	var key []byte
	if keyFile != "" {
		key, err = ioutil.ReadFile(keyFile)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Printf("\n----WARNING----\n")
		fmt.Printf("NO KEYFILE GIVEN, PAYLOAD WILL NOT BE ENCRYPTED\n")
		fmt.Printf("\nARE YOU SURE YOU WANT TO CONTINUE? (y/N): ")

		reader := bufio.NewReader(os.Stdin)
		confirm, _ := reader.ReadString('\n')
		confirm = strings.Trim(confirm, "\n")
		switch confirm {
		case "n", "no", "N", "NO", "No":
			os.Exit(0)
		case "y", "yes", "Y", "YES", "Yes":
			fmt.Printf("Proceeding with no encryption.\n")
			fmt.Printf("----WARNING END----\n")
			break
		default:
			os.Exit(0)
		}

		key = nil
	}

	// Encrypt our payload if needed, and create out final payload string
	var finalPayloadString string
	if key != nil {
		finalPayloadString = payload.EncryptAES(key, payloadString)
	} else {
		finalPayloadString = payloadString
	}

	payload := payload.BuildPayload(finalPayloadString, "puNK", shuffle)
	burnedChunks := injection.Inject(referenceChunks, payload, shuffle)
	// Export to file?
	if target != "" {
		injection.Rebuild(burnedChunks, target)
		fmt.Printf("\n----WRITTEN CHUNKS----\n")
		injection.PrintChunks(burnedChunks)
		fmt.Printf("----END WRITTEN CHUNKS----\n")
	} else {
		fmt.Printf("\nNo target. Would've written the following bytes;\n")
		injection.PrintChunks(burnedChunks)
	}
}
