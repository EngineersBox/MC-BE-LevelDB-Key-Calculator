package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/chunk"
)

func main() {
	xCoord := flag.Int("x", 0, "X Coordinate")
	yCoord := flag.Int("y", 0, "Y Coordinate")
	zCoord := flag.Int("z", 0, "Z Coordinate")
	worldType := flag.String("type", "overworld", "<overworld | nether | end>")

	flag.Parse()

	var levelDBKey strings.Builder
	// Append the X chunk coord
	chunk.ChunkCoordLittleEndian(&levelDBKey, *xCoord, chunk.ChunkSizeX)
	// Append the Y chunk coord
	chunk.ChunkCoordLittleEndian(&levelDBKey, *zCoord, chunk.ChunkSizeZ)

	// Append dimension keys if nether or end
	if *worldType == "nether" {
		levelDBKey.WriteString("FFFFFFFF")
	} else if *worldType == "end" {
		levelDBKey.WriteString("01000000")
	}

	// Add the subchunk prefix (47 = 0x2f)
	levelDBKey.WriteString(fmt.Sprintf("%x", chunk.SubChunkPrefixTag))

	yChunk := *yCoord / chunk.SubChunkSizeY
	if yChunk < 10 {
		// If the value is less than 10, add a 0 to ensure
		// the string hex value is prefixed properly
		levelDBKey.WriteString("0")
	}
	// Append the Y subchunk coord
	levelDBKey.WriteString(fmt.Sprintf("%x", int8(yChunk)))

	fmt.Println(levelDBKey.String())
}
