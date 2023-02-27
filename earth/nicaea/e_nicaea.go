package nicaea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NicaeaEmpire interface {
    feud.Empire
}

type 尼西亚帝国NicaeaEmpire struct {
	feud.BaseEmpire
}

var ENicaea尼西亚帝国 NicaeaEmpire = &尼西亚帝国NicaeaEmpire{}

func init() {
	f := ENicaea尼西亚帝国.(*尼西亚帝国NicaeaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "nicaea",
		TitleName: "尼西亚帝国",
		TitleCode: "e_nicaea",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
