package zemigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯帕尔内内SparneneBarony struct {
	feud.BaseBarony
}

var BSparnene斯帕尔内内 feud.Barony = &斯帕尔内内SparneneBarony{}

func init() {
    f := BSparnene斯帕尔内内.(*斯帕尔内内SparneneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sparnene",
		TitleName: "斯帕尔内内",
		TitleCode: "b_sparnene",
	}
}
