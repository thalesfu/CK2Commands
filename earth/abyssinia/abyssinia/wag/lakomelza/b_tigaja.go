package lakomelza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提嘎加TigajaBarony struct {
	feud.BaseBarony
}

var BTigaja提嘎加 feud.Barony = &提嘎加TigajaBarony{}

func init() {
	f := BTigaja提嘎加.(*提嘎加TigajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tigaja",
		TitleName: "提嘎加",
		TitleCode: "b_tigaja",
	}
}
