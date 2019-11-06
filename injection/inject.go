package injection

func Inject(chunks []Chunk, payload Chunk) []Chunk {
	var burnedChunks []Chunk
	//var IDATpassed bool

	for i := range chunks {
		//if IDATpassed {
		if i == 2 {
			burnedChunks = append(burnedChunks, payload)
			//IDATpassed = false
		}
		burnedChunks = append(burnedChunks, chunks[i])

		// Check if this is the IDAT chunk
		if string(chunks[i].CType) == "IDAT" {
			//IDATpassed = true
		}
	}
	return burnedChunks
}
