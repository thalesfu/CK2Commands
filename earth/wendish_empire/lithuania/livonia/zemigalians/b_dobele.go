package zemigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多贝莱DobeleBarony struct {
	feud.BaseBarony
}

var BDobele多贝莱 feud.Barony = &多贝莱DobeleBarony{}

func init() {
    f := BDobele多贝莱.(*多贝莱DobeleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dobele",
		TitleName: "多贝莱",
		TitleCode: "b_dobele",
	}
}
