package zachlumia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯通StonBarony struct {
	feud.BaseBarony
}

var BSton斯通 feud.Barony = &斯通StonBarony{}

func init() {
    f := BSton斯通.(*斯通StonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ston",
		TitleName: "斯通",
		TitleCode: "b_ston",
	}
}
