package latin_empire

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Latin_empireEmpire interface {
    feud.Empire
}

type 拉丁帝国Latin_empireEmpire struct {
	feud.BaseEmpire
}

var ELatin_empire拉丁帝国 Latin_empireEmpire = &拉丁帝国Latin_empireEmpire{}

func init() {
	f := ELatin_empire拉丁帝国.(*拉丁帝国Latin_empireEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "latin_empire",
		TitleName: "拉丁帝国",
		TitleCode: "e_latin_empire",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
