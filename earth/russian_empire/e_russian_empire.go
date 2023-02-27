package russian_empire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Russian_empireEmpire interface {
    feud.Empire
}

type 俄罗斯Russian_empireEmpire struct {
	feud.BaseEmpire
}

var ERussian_empire俄罗斯 Russian_empireEmpire = &俄罗斯Russian_empireEmpire{}

func init() {
	f := ERussian_empire俄罗斯.(*俄罗斯Russian_empireEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "russian_empire",
		TitleName: "俄罗斯",
		TitleCode: "e_russian_empire",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
