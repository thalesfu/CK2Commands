package china_west_governor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type China_west_governorEmpire interface {
    feud.Empire
}

type 西域都护府China_west_governorEmpire struct {
	feud.BaseEmpire
}

var EChina_west_governor西域都护府 China_west_governorEmpire = &西域都护府China_west_governorEmpire{}

func init() {
	f := EChina_west_governor西域都护府.(*西域都护府China_west_governorEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "china_west_governor",
		TitleName: "西域都护府",
		TitleCode: "e_china_west_governor",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
