package chunks

type TagType int

const (
	ChunkVersion        TagType = 0x2c
	Data2D              TagType = 0x2d
	Data2DLegacy        TagType = 0x2e
	SubChunkPrefix      TagType = 0x2f
	LegacyTerrain       TagType = 0x30
	BlockEntity         TagType = 0x31
	Entity              TagType = 0x32
	PendingTicks        TagType = 0x33
	BlockExtraData      TagType = 0x35
	BiomeState          TagType = 0x35
	UNUSED              TagType = 0x37
	BorderBlocks        TagType = 0x38
	HardCodedSpawnAreas TagType = 0x39
	RandomTicks         TagType = 0x3a
	Checksums           TagType = 0x3b
	ChunkVersionV116100 TagType = 0x76
)

var TagTypes = map[string]TagType{
	"ChunkVersion":        ChunkVersion,
	"Data2D":              Data2D,
	"Data2DLegacy":        Data2DLegacy,
	"SubChunkPrefix":      SubChunkPrefix,
	"LegacyTerrain":       LegacyTerrain,
	"BlockEntity":         BlockEntity,
	"Entity":              Entity,
	"PendingTicks":        PendingTicks,
	"BlockExtraData":      BlockExtraData,
	"BiomeState":          BiomeState,
	"UNUSED":              UNUSED,
	"BorderBlocks":        BorderBlocks,
	"HardCodedSpawnAreas": HardCodedSpawnAreas,
	"RandomTicks":         RandomTicks,
	"Checksums":           Checksums,
	"ChunkVersionV116100": ChunkVersionV116100,
}
