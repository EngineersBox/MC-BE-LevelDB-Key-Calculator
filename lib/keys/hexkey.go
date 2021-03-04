package lib

import (
	"fmt"
	"strings"

	tagbytes "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/chunks"
	world "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/world"
)

type HexKey struct {
	ChunkX   int32
	ChunkY   int8
	ChunkZ   int32
	WorldKey world.WorldType
	TagKey   tagbytes.TagType
}

func (hk *HexKey) ToString() string {
	var levelDBKey strings.Builder

	// Append the X subchunk coord
	levelDBKey.WriteString(fmt.Sprintf("%x", hk.ChunkX))

	// Append the Z subchunk coord
	levelDBKey.WriteString(fmt.Sprintf("%x", hk.ChunkZ))

	// Append dimension keys if nether or end
	levelDBKey.WriteString(string(hk.WorldKey))

	// Add the tag key
	levelDBKey.WriteString(fmt.Sprintf("%x", hk.TagKey))

	if hk.ChunkY < 10 {
		// If the value is less than 10, add a 0 to ensure
		// the string hex value is prefixed properly
		levelDBKey.WriteString("0")
	}
	// Append the Y subchunk coord
	levelDBKey.WriteString(fmt.Sprintf("%x", hk.ChunkY))

	return levelDBKey.String()
}
