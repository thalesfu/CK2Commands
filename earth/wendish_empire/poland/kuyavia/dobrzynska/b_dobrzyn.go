package dobrzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多布任DobrzynBarony struct {
	feud.BaseBarony
}

var BDobrzyn多布任 feud.Barony = &多布任DobrzynBarony{}

func init() {
    f := BDobrzyn多布任.(*多布任DobrzynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobrzyn",
		TitleName: "多布任",
		TitleCode: "b_dobrzyn",
	}
}
