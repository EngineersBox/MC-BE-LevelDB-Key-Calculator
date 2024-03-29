# Minecraft Bedrock Edition LevelDB Key Calculator
A calculator to generate the hex keys for retrieving data from a Minecraft: Bedrock Edition LDB (chunk data) database

## Overview

This is a very simple tool used to convert world coordinates into chunk hex keys. Chunk hex keys are used to access sections of the LDB containing chunk data.

## Using the Tool

First step is to clone the repo and `cd` into it, go ahead and do that and come back here after.

### Command Options

| Option  | Type      | Default Value    | Valid Values                                                                                                                                                                                                                                                                           |
|---------|-----------|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `-x`    | `int32`   | `0`              | Any valid integer value                                                                                                                                                                                                                                                                |
| `-y`    | `int32`   | `0`              | Any valid integer value                                                                                                                                                                                                                                                                |
| `-z`    | `int32`   | `0`              | Any valid integer value                                                                                                                                                                                                                                                                |
| `-chunkCoords`    | `bool`   | `false` | Any valid boolean value                                                                                                                                                                                                                                                                |
| `-type` | `string`  | `overworld`      | `overworld`<br>`nether`<br>`end`                                                                                                                                                                                                                                                       |
| `-tag`  | `TagType` | `SubChunkPrefix` | `ChunkVersion`<br>`Data2D`<br>`Data2DLegacy`<br>`SubChunkPrefix`<br>`LegacyTerrain`<br>`BlockEntity`<br>`Entity`<br>`PendingTicks`<br>`BlockExtraData`<br>`BiomeState`<br>`UNUSED`<br>`BorderBlocks`<br>`HardCodedSpawnAreas`<br>`RandomTicks`<br>`Checksums`<br>`ChunkVersionV116100` |


### Command Syntax

```shell
go run main.go hexkey [OPTIONS]
```

For example, running the following:

```shell
go run main.go hexkey -x 413 -z 54 -y 105 -type nether -tag PendingTicks
```

produces:

```yaml
190000003000000ffffffff3306
```

## Hex Key Format

Minecraft Gamepedia Link: <https://minecraft.gamepedia.com/Bedrock_Edition_level_format>

```golang
<LE Chunk X Coord><LE Chunk Z Coord>[<Nether Key | End Key>]<Tag Byte><BE SubChunk Y Coord>
```

For example, using the coordinates `X: 413, Y: 105, Z: 54`, the coresponding chunk key is (split by section):

```yaml
19000000 03000000 2f 06
```

Another example, using the coordinates `X: 110, Y: 2, Z: -22` with chunk coord format, the coresponding chunk key is (split by section):

```yaml
6e000000 eaffffff 2f 02
```

*Legend:*
| Key          | Expansion     |
|--------------|---------------|
| `LE`         | Little Endian |
| `BE`         | Big Endian    |
| `Nether Key` | `0xffffffff`  |
| `End Key`    | `0x01000000`  |

### Manually Calculating a Hex Key

X, Y, Z is the typical coordinate order, but when dealing with the data we find X, Z, Y is the order of greatest-to-least significance, which is why this explanation tends to express X, Z, Y ordering.

All division below is of course integer division. The remainder/modulus will be used to find the byte offset within the subchunk data. X, Z, and dimension are 32-bit signed integers in little endian byte order. In the examples below, I've bolded the chunk Z coordinate for clarity.

Each chunk is `16x16x256` (X, Z, Y), and the subchunk block data keys are `16` high. So for x, z, y coordinates of `413, 54, 105`:

- chunk X = `413 / 16 = 25` or `0x19` signed 32-bit integer in little endian byte order (`[0x19,0, 0, 0] == 19000000`)
- chunk Z = `54 / 16 = 3` (`[0x3, 0, 0, 0] == 03000000`) 

So all keys beginning with `1900000003000000` are about this coordinate's chunk. (In the overworld; other dimensions add a 32-bit dimension ID, so the same coordinates in the Nether I think have keys that start with `1900000003000000ffffffff` and `190000000300000001000000` for the End.)

The tags and subchunk indexes are 8-bit values. (Unsigned? Not sure it matters as there are no negative Y chunk coordinates and no tags < `0` or > `127`.)

`47` (`[0x2F]`) is the subchunk prefix tag, so all keys beginning with `19000000030000002f` are the Y subchunks for this coordinate.

- subchunk Y = `105 / 16 = 6` (`[0x06]`)

So, the subchunk key for X=`413`, Z=`54`, Y=`105` is `19000000030000002f06`

*Attribution*: This section was taken from the README.md on the McpeTool repo, see here for the full document: <https://github.com/midnightfreddie/McpeTool/tree/master/docs#how-to-convert-world-coordinates-to-leveldb-keys>
