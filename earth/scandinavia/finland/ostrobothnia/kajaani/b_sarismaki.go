package kajaani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨里斯马基SarismakiBarony struct {
	feud.BaseBarony
}

var BSarismaki萨里斯马基 feud.Barony = &萨里斯马基SarismakiBarony{}

func init() {
	f := BSarismaki萨里斯马基.(*萨里斯马基SarismakiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarismaki",
		TitleName: "萨里斯马基",
		TitleCode: "b_sarismaki",
	}
}
