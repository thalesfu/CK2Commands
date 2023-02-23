package timurids

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TimuridsEmpire interface {
	feud.Empire
}

type 帖木儿TimuridsEmpire struct {
	feud.BaseEmpire
}

var ETimurids帖木儿 TimuridsEmpire = &帖木儿TimuridsEmpire{}

func init() {
	f := ETimurids帖木儿.(*帖木儿TimuridsEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "timurids",
		TitleName: "帖木儿",
		TitleCode: "e_timurids",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
