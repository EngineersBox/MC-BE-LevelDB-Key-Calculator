package chunks

import (
	"encoding/binary"

	bytearrays "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/bytearrays"
)

const (
	ChunkSizeX    = 16
	ChunkSizeY    = 256
	ChunkSizeZ    = 16
	SubChunkSizeY = 16
)

func ChunkCoordLittleEndian(coord int, chunkSize int, chunkCoords bool) uint32 {
	// Convert an int32 into big endian byte array
	var baseCoord int32
	if chunkCoords {
		baseCoord = int32(coord)
	} else {
		baseCoord = int32(coord / chunkSize)
	}
	var chunkBytes []byte
	if baseCoord >= 0 {
		chunkBytes = bytearrays.IntToByteArray(baseCoord)
	} else {
		chunkBytes = bytearrays.TwosComplement(baseCoord)
	}
	// Create an int32 from the little endian byte array
	chunk := binary.LittleEndian.Uint32(chunkBytes)
	// Append the little endian int32
	return chunk
}
