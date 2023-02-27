package kaisereia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨勒兹SarizBarony struct {
	feud.BaseBarony
}

var BSariz萨勒兹 feud.Barony = &萨勒兹SarizBarony{}

func init() {
    f := BSariz萨勒兹.(*萨勒兹SarizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sariz",
		TitleName: "萨勒兹",
		TitleCode: "b_sariz",
	}
}
