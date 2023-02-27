package troyes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克莱尔沃ClairvauxBarony struct {
	feud.BaseBarony
}

var BClairvaux克莱尔沃 feud.Barony = &克莱尔沃ClairvauxBarony{}

func init() {
    f := BClairvaux克莱尔沃.(*克莱尔沃ClairvauxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clairvaux",
		TitleName: "克莱尔沃",
		TitleCode: "b_clairvaux",
	}
}
