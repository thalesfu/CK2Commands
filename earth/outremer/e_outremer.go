package outremer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OutremerEmpire interface {
	feud.Empire
}

type 海外OutremerEmpire struct {
	feud.BaseEmpire
}

var EOutremer海外 OutremerEmpire = &海外OutremerEmpire{}

func init() {
	f := EOutremer海外.(*海外OutremerEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "outremer",
		TitleName: "海外",
		TitleCode: "e_outremer",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
