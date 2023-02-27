package pirates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PiratesEmpire interface {
    feud.Empire
}

type 海盗PiratesEmpire struct {
	feud.BaseEmpire
}

var EPirates海盗 PiratesEmpire = &海盗PiratesEmpire{}

func init() {
	f := EPirates海盗.(*海盗PiratesEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "pirates",
		TitleName: "海盗",
		TitleCode: "e_pirates",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
