package lib

import (
	"fmt"

	chunk "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/chunks"
	tagbytes "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/chunks"
	world "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/world"
)

// Coordinates ... Stores the (X, Y, Z) world location
type Coordinates struct {
	X *int
	Y *int
	Z *int
}

// Attributes ... Stores the required world and chunk tag attributes
type Attributes struct {
	WorldType *string
	TagType   *string
}

// LDBKeyParameters .. Stores all required parameters to create the hex key
type LDBKeyParameters struct {
	Coords Coordinates
	Attrs  Attributes
}

func (lkp *LDBKeyParameters) CalculateHexKey() (hexKey HexKey) {
	// Append the X chunk coord
	hexKey.ChunkX = chunk.ChunkCoordLittleEndian(*lkp.Coords.X, chunk.ChunkSizeX)
	// Append the Z chunk coord
	hexKey.ChunkZ = chunk.ChunkCoordLittleEndian(*lkp.Coords.Z, chunk.ChunkSizeZ)

	// Append dimension keys if nether or end
	enumWorldTypeValue, ok := world.WorldTypes[*lkp.Attrs.WorldType]
	if !ok {
		panic(fmt.Sprintf("No such world type: %s", *lkp.Attrs.WorldType))
	}
	hexKey.WorldKey = enumWorldTypeValue

	// Add the tag key
	enumTagTypeValue, ok := tagbytes.TagTypes[*lkp.Attrs.TagType]
	if !ok {
		panic(fmt.Sprintf("No such tag type: %s", *lkp.Attrs.TagType))
	}
	hexKey.TagKey = enumTagTypeValue

	// Append the Y chunk coord
	hexKey.ChunkY = int8(*lkp.Coords.Y / chunk.SubChunkSizeY)

	return hexKey
}
