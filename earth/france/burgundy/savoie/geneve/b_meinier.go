package geneve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅涅MeinierBarony struct {
	feud.BaseBarony
}

var BMeinier梅涅 feud.Barony = &梅涅MeinierBarony{}

func init() {
	f := BMeinier梅涅.(*梅涅MeinierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meinier",
		TitleName: "梅涅",
		TitleCode: "b_meinier",
	}
}
