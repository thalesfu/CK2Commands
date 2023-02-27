package innse_gall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯利特SleatBarony struct {
	feud.BaseBarony
}

var BSleat斯利特 feud.Barony = &斯利特SleatBarony{}

func init() {
    f := BSleat斯利特.(*斯利特SleatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sleat",
		TitleName: "斯利特",
		TitleCode: "b_sleat",
	}
}
