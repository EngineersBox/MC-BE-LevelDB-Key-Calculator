import (
	"ByteArrays"
)

const (
	ChunkSizeX        = 16
	SubChunkSizeY     = 16
	ChunkSizeZ        = 16
	SubChunkPrefixTag = 0x2f
)

func ChunkCoordLittleEndian(levelDBKey *strings.Builder, coord int, chunkSize int) {
	// Convert an int32 into big endian byte array
	chunkBytes := ByteArrays.IntToByteArray(int32(coord / chunkSize))
	// Reverse the big endian byte array to get little endian format
	ByteArrays.reverseAny(chunkBytes)
	// Create an int32 from the little endian byte array
	chunk := ByteArrays.ByteArrayToInt(chunkBytes)
	// Append the little endian int32
	(*levelDBKey).WriteString(fmt.Sprintf("%x", chunk))
}