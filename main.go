package main

import (
	"flag"
	"fmt"

	lkp "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/keys"
)

func main() {
	var keyParameters lkp.LDBKeyParameters = lkp.LDBKeyParameters{
		Coords: lkp.Coordinates{
			X: flag.Int("x", 0, "-x <int>"),
			Y: flag.Int("y", 0, "-y <int>"),
			Z: flag.Int("z", 0, "-z <int>"),
		},
		Attrs: lkp.Attributes{
			WorldType: flag.String("type", "Overworld", "-type <overworld | nether | end>"),
			TagType:   flag.String("tag", "SubChunkPrefix", "-tag <TagType>"),
		},
	}

	flag.Parse()

	var hexKey lkp.HexKey = keyParameters.CalculateHexKey()

	fmt.Println(hexKey.ToString())

}
