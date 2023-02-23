package rutog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 果巴卡GobakBarony struct {
	feud.BaseBarony
}

var BGobak果巴卡 feud.Barony = &果巴卡GobakBarony{}

func init() {
	f := BGobak果巴卡.(*果巴卡GobakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gobak",
		TitleName: "果巴卡",
		TitleCode: "b_gobak",
	}
}
