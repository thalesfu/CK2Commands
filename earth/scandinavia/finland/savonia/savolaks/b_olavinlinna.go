package savolaks

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥拉维堡OlavinlinnaBarony struct {
	feud.BaseBarony
}

var BOlavinlinna奥拉维堡 feud.Barony = &奥拉维堡OlavinlinnaBarony{}

func init() {
    f := BOlavinlinna奥拉维堡.(*奥拉维堡OlavinlinnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olavinlinna",
		TitleName: "奥拉维堡",
		TitleCode: "b_olavinlinna",
	}
}
