package viviers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维维耶ViviersBarony struct {
	feud.BaseBarony
}

var BViviers维维耶 feud.Barony = &维维耶ViviersBarony{}

func init() {
    f := BViviers维维耶.(*维维耶ViviersBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "viviers",
		TitleName: "维维耶",
		TitleCode: "b_viviers",
	}
}
