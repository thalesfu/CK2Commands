package chagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChagataiEmpire interface {
	feud.Empire
}

type 察合台ChagataiEmpire struct {
	feud.BaseEmpire
}

var EChagatai察合台 ChagataiEmpire = &察合台ChagataiEmpire{}

func init() {
	f := EChagatai察合台.(*察合台ChagataiEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "chagatai",
		TitleName: "察合台",
		TitleCode: "e_chagatai",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
