package shiite

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShiiteEmpire interface {
	feud.Empire
}

type 什叶派哈里发国ShiiteEmpire struct {
	feud.BaseEmpire
}

var EShiite什叶派哈里发国 ShiiteEmpire = &什叶派哈里发国ShiiteEmpire{}

func init() {
	f := EShiite什叶派哈里发国.(*什叶派哈里发国ShiiteEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "shiite",
		TitleName: "什叶派哈里发国",
		TitleCode: "e_shiite",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
