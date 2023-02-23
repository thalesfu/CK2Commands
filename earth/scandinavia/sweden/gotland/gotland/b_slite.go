package gotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯利特SliteBarony struct {
	feud.BaseBarony
}

var BSlite斯利特 feud.Barony = &斯利特SliteBarony{}

func init() {
	f := BSlite斯利特.(*斯利特SliteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slite",
		TitleName: "斯利特",
		TitleCode: "b_slite",
	}
}
