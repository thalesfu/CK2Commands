package keriya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 车江ChejiangBarony struct {
	feud.BaseBarony
}

var BChejiang车江 feud.Barony = &车江ChejiangBarony{}

func init() {
	f := BChejiang车江.(*车江ChejiangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chejiang",
		TitleName: "车江",
		TitleCode: "b_chejiang",
	}
}
