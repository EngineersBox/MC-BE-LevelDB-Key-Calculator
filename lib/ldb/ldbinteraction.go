package ldb

import (
	lkp "github.com/EngineersBox/MC-BE-LevelDB-Key-Calculator/lib/keys"
)

type LDBInteraction struct {
	InteractionMethod *string
	Port              *int
	WorldPath         *string
	LDBEntryPath      *string
	OutFile           *string
	Parameters        lkp.LDBKeyParameters
}
