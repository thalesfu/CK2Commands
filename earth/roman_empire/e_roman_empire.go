package roman_empire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Roman_empireEmpire interface {
    feud.Empire
}

type 罗马帝国Roman_empireEmpire struct {
	feud.BaseEmpire
}

var ERoman_empire罗马帝国 Roman_empireEmpire = &罗马帝国Roman_empireEmpire{}

func init() {
	f := ERoman_empire罗马帝国.(*罗马帝国Roman_empireEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "roman_empire",
		TitleName: "罗马帝国",
		TitleCode: "e_roman_empire",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
