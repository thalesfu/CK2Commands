package bulgaria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BulgariaEmpire interface {
	feud.Empire
}

type 保加利亚BulgariaEmpire struct {
	feud.BaseEmpire
}

var EBulgaria保加利亚 BulgariaEmpire = &保加利亚BulgariaEmpire{}

func init() {
	f := EBulgaria保加利亚.(*保加利亚BulgariaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "bulgaria",
		TitleName: "保加利亚",
		TitleCode: "e_bulgaria",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
