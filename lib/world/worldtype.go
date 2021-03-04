package lib

type WorldType string

const (
	Overworld WorldType = ""
	Nether    WorldType = "ffffffff"
	End       WorldType = "01000000"
)

var WorldTypes = map[string]WorldType{
	"Overworld": Overworld,
	"Nether":    Nether,
	"End":       End,
}
