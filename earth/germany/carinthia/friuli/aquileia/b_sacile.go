package aquileia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨奇莱SacileBarony struct {
	feud.BaseBarony
}

var BSacile萨奇莱 feud.Barony = &萨奇莱SacileBarony{}

func init() {
	f := BSacile萨奇莱.(*萨奇莱SacileBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sacile",
		TitleName: "萨奇莱",
		TitleCode: "b_sacile",
	}
}
