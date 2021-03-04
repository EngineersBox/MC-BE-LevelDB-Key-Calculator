package main

import (
	"flag"
	"fmt"
	"strings"

	chunk "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/chunks"
	tagbytes "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/chunks"
	world "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/world"
)

func main() {
	var xCoord *int = flag.Int("x", 0, "-x <int>")
	var yCoord *int = flag.Int("y", 0, "-y <int>")
	var zCoord *int = flag.Int("z", 0, "-z <int>")
	var worldType *string = flag.String("type", "Overworld", "-type <overworld | nether | end>")
	var tagType *string = flag.String("tag", "SubChunkPrefix", "-tag <TagType>")

	flag.Parse()

	var levelDBKey strings.Builder

	// Append the X chunk coord
	chunk.ChunkCoordLittleEndian(&levelDBKey, *xCoord, chunk.ChunkSizeX)
	// Append the Y chunk coord
	chunk.ChunkCoordLittleEndian(&levelDBKey, *zCoord, chunk.ChunkSizeZ)

	// Append dimension keys if nether or end
	enumWorldTypeValue, ok := world.WorldTypes[*worldType]
	if !ok {
		panic(fmt.Sprintf("No such world type: %s", *worldType))
	}
	levelDBKey.WriteString(string(enumWorldTypeValue))

	// Add the subchunk prefix (47 = 0x2f)
	enumTagTypeValue, ok := tagbytes.TagTypes[*tagType]
	if !ok {
		panic(fmt.Sprintf("No such tag type: %s", *tagType))
	}
	levelDBKey.WriteString(fmt.Sprintf("%x", enumTagTypeValue))

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
