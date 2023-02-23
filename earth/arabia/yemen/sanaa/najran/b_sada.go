package najran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨达SadaBarony struct {
	feud.BaseBarony
}

var BSada萨达 feud.Barony = &萨达SadaBarony{}

func init() {
	f := BSada萨达.(*萨达SadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sada",
		TitleName: "萨达",
		TitleCode: "b_sada",
	}
}
