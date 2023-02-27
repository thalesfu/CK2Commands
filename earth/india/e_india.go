package india

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IndiaEmpire interface {
    feud.Empire
}

type 印度IndiaEmpire struct {
	feud.BaseEmpire
}

var EIndia印度 IndiaEmpire = &印度IndiaEmpire{}

func init() {
	f := EIndia印度.(*印度IndiaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "india",
		TitleName: "印度",
		TitleCode: "e_india",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
