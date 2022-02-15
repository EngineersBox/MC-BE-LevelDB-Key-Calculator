package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	keys "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/keys"
	ldb "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/ldb"
	"github.com/midnightfreddie/McpeTool/api"
	"github.com/midnightfreddie/McpeTool/world"
)

// TODO: Move all the stuff in this file to relevant modules

func NewMcpeToolApi(interaction ldb.LDBInteraction) {
	world, err := world.OpenWorld(*interaction.WorldPath)
	if err != nil {
		panic(err)
	}
	defer world.Close()
	fmt.Println("Starting API server for world at " + *interaction.WorldPath)
	fmt.Println("REST at http://localhost:" + fmt.Sprintf("%d", *interaction.Port) + "/api/v1/db")
	err = api.Server(&world, "localhost", fmt.Sprintf("%d", *interaction.Port))
	if err != nil {
		panic(err)
	}
}

type LdbAPIResponse struct {
	ApiVersion string
	HexKey     string
	Base64Data string
	Nbt2Json   map[string]interface{}
}

type LdbKeysResponse struct {
	ApiVersion string
	Keys       []struct {
		HexKey string
	}
}

const (
	OutFilePerm fs.FileMode = 0644
)

func main() {
	hexKeyCommand := flag.NewFlagSet("hexkey", flag.ExitOnError)
	ldbCommand := flag.NewFlagSet("ldb", flag.ExitOnError)

	var keyParameters keys.LDBKeyParameters = keys.LDBKeyParameters{
		Coords: keys.Coordinates{
			X: hexKeyCommand.Int("x", 0, "-x <int>"),
			Y: hexKeyCommand.Int("y", 0, "-y <int>"),
			Z: hexKeyCommand.Int("z", 0, "-z <int>"),
		},
		Attrs: keys.Attributes{
			WorldType:   hexKeyCommand.String("type", "Overworld", "-type <overworld | nether | end>"),
			TagType:     hexKeyCommand.String("tag", "SubChunkPrefix", "-tag <TagType>"),
			ChunkCoords: hexKeyCommand.Bool("chunkCoords", false, "-chunkCoords <true | false>"),
		},
	}

	var ldbParameters ldb.LDBInteraction = ldb.LDBInteraction{
		InteractionMethod: ldbCommand.String("method", "GET", "-method <LIST | GET | PUT | DELETE>"),
		Port:              ldbCommand.Int("port", 8090, "-port <int>"),
		WorldPath:         ldbCommand.String("path", "", "-path <world path>"),
		LDBEntryPath:      ldbCommand.String("entry", "", "-entry <path/to/ldb/entry/body>"),
		OutFile:           ldbCommand.String("out", "out.json", "-out <filepath>"),
		Parameters: keys.LDBKeyParameters{
			Coords: keys.Coordinates{
				X: ldbCommand.Int("x", 0, "-x <int>"),
				Y: ldbCommand.Int("y", 0, "-y <int>"),
				Z: ldbCommand.Int("z", 0, "-z <int>"),
			},
			Attrs: keys.Attributes{
				WorldType: ldbCommand.String("type", "Overworld", "-type <overworld | nether | end>"),
				TagType:   ldbCommand.String("tag", "SubChunkPrefix", "-tag <TagType>"),
			},
		},
	}

	switch os.Args[1] {
	case "hexkey":
		hexKeyCommand.Parse(os.Args[2:])
	case "ldb":
		ldbCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if hexKeyCommand.Parsed() {
		// Required Flags
		if keyParameters.Coords.X == nil ||
			keyParameters.Coords.Y == nil ||
			keyParameters.Coords.Z == nil {
			hexKeyCommand.PrintDefaults()
			os.Exit(1)
		}

		var hexKey keys.HexKey = keyParameters.CalculateHexKey()
		fmt.Println(hexKey.ToString())
	}

	if ldbCommand.Parsed() {
		// Required Flags
		if ldbParameters.Parameters.Coords.X == nil ||
			ldbParameters.Parameters.Coords.Y == nil ||
			ldbParameters.Parameters.Coords.Z == nil ||
			ldbParameters.InteractionMethod == nil ||
			ldbParameters.Port == nil ||
			ldbParameters.WorldPath == nil {
			hexKeyCommand.PrintDefaults()
			os.Exit(1)
		}

		var hexKey keys.HexKey = ldbParameters.Parameters.CalculateHexKey()
		go NewMcpeToolApi(ldbParameters)
		time.Sleep(1)
		if *ldbParameters.InteractionMethod == "GET" {
			LdbGetEndpoint(ldbParameters, hexKey)
		} else if *ldbParameters.InteractionMethod == "LIST" {
			LdbListEndpoint(ldbParameters, hexKey)
		} else if *ldbParameters.InteractionMethod == "PUT" {
			LdbPutEndpoint(ldbParameters, hexKey)
		} else if *ldbParameters.InteractionMethod == "DELETE" {
			LdbDeleteEndpoint(ldbParameters, hexKey)
		}
	}

}

func LdbGetEndpoint(ldbParameters ldb.LDBInteraction, hexKey keys.HexKey) {
	resp, err := http.Get(
		fmt.Sprintf(
			"http://localhost:%d/api/v1/db/%s?json",
			*ldbParameters.Port,
			hexKey.ToString(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("non 2XX API response")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyStruct := LdbAPIResponse{}
	json.Unmarshal(bodyBytes, &bodyStruct)
	if ldbParameters.OutFile == nil {
		nbtJson, err := json.Marshal(bodyStruct.Nbt2Json)
		if err != nil {
			panic(err)
		}
		fmt.Println(nbtJson)
		return
	}
	nbtJson, err := json.Marshal(bodyStruct.Nbt2Json)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(*ldbParameters.OutFile, []byte(nbtJson), OutFilePerm)
	if err != nil {
		fmt.Println("could not write to file:", err)
		fmt.Println(nbtJson)
	}
}

func LdbListEndpoint(ldbParameters ldb.LDBInteraction, hexKey keys.HexKey) {
	fmt.Println("Pulling all LDB keys... this can take a while")
	resp, err := http.Get(
		fmt.Sprintf(
			"http://localhost:%d/api/v1/db",
			*ldbParameters.Port,
		),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("non 2XX API response")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	keysStruct := LdbKeysResponse{}
	json.Unmarshal(bodyBytes, &keysStruct)
	if ldbParameters.OutFile == nil {
		keyJson, err := json.Marshal(keysStruct.Keys)
		if err != nil {
			panic(err)
		}
		fmt.Println(keyJson)
		return
	}
	keyJson, err := json.Marshal(keysStruct.Keys)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(*ldbParameters.OutFile, []byte(keyJson), OutFilePerm)
	if err != nil {
		fmt.Println("could not write to file:", err)
		fmt.Println(keyJson)
	}
}

func LdbPutEndpoint(ldbParameters ldb.LDBInteraction, hexKey keys.HexKey) {
	ldbEntryBody, err := ioutil.ReadFile(*ldbParameters.LDBEntryPath)
	_, err = http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"http://localhost:%d/api/v1/db/%s?json",
			*ldbParameters.Port,
			hexKey.ToString(),
		),
		bytes.NewBuffer(ldbEntryBody),
	)
	if err != nil {
		panic(err)
	}
}

func LdbDeleteEndpoint(ldbParameters ldb.LDBInteraction, hexKey keys.HexKey) {
	_, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"http://localhost:%d/api/v1/db/%s",
			*ldbParameters.Port,
			hexKey.ToString(),
		),
		nil,
	)
	if err != nil {
		panic(err)
	}
}
