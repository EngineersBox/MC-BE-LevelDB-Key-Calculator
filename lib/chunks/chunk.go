package chunks

import (
	bytearrays "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/bytearrays"
)

const (
	ChunkSizeX    = 16
	ChunkSizeY    = 256
	ChunkSizeZ    = 16
	SubChunkSizeY = 16
)

func ChunkCoordLittleEndian(coord int, chunkSize int) int32 {
	// Convert an int32 into big endian byte array
	chunkBytes := bytearrays.IntToByteArray(int32(coord / chunkSize))
	// Reverse the big endian byte array to get little endian format
	bytearrays.ReverseAny(chunkBytes)
	// Create an int32 from the little endian byte array
	chunk := bytearrays.ByteArrayToInt(chunkBytes)
	// Append the little endian int32
	// (*levelDBKey).WriteString(fmt.Sprintf("%x", chunk))
	return chunk
}
