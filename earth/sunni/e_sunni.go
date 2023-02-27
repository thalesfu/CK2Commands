package sunni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SunniEmpire interface {
    feud.Empire
}

type 逊尼派哈里发国SunniEmpire struct {
	feud.BaseEmpire
}

var ESunni逊尼派哈里发国 SunniEmpire = &逊尼派哈里发国SunniEmpire{}

func init() {
	f := ESunni逊尼派哈里发国.(*逊尼派哈里发国SunniEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "sunni",
		TitleName: "逊尼派哈里发国",
		TitleCode: "e_sunni",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
