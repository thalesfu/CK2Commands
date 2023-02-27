package cordoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CordobaEmpire interface {
    feud.Empire
}

type 科尔多瓦CordobaEmpire struct {
	feud.BaseEmpire
}

var ECordoba科尔多瓦 CordobaEmpire = &科尔多瓦CordobaEmpire{}

func init() {
	f := ECordoba科尔多瓦.(*科尔多瓦CordobaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "cordoba",
		TitleName: "科尔多瓦",
		TitleCode: "e_cordoba",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
