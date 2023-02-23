package chieti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰恰诺LancianoBarony struct {
	feud.BaseBarony
}

var BLanciano兰恰诺 feud.Barony = &兰恰诺LancianoBarony{}

func init() {
	f := BLanciano兰恰诺.(*兰恰诺LancianoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lanciano",
		TitleName: "兰恰诺",
		TitleCode: "b_lanciano",
	}
}
