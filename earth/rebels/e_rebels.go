package rebels

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RebelsEmpire interface {
    feud.Empire
}

type 叛军RebelsEmpire struct {
	feud.BaseEmpire
}

var ERebels叛军 RebelsEmpire = &叛军RebelsEmpire{}

func init() {
	f := ERebels叛军.(*叛军RebelsEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "rebels",
		TitleName: "叛军",
		TitleCode: "e_rebels",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
