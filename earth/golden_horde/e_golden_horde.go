package golden_horde

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Golden_hordeEmpire interface {
    feud.Empire
}

type 金帐汗国Golden_hordeEmpire struct {
	feud.BaseEmpire
}

var EGolden_horde金帐汗国 Golden_hordeEmpire = &金帐汗国Golden_hordeEmpire{}

func init() {
	f := EGolden_horde金帐汗国.(*金帐汗国Golden_hordeEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "golden_horde",
		TitleName: "金帐汗国",
		TitleCode: "e_golden_horde",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
