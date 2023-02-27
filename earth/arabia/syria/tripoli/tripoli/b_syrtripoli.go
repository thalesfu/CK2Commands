package tripoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 的黎波里SyrtripoliBarony struct {
	feud.BaseBarony
}

var BSyrtripoli的黎波里 feud.Barony = &的黎波里SyrtripoliBarony{}

func init() {
    f := BSyrtripoli的黎波里.(*的黎波里SyrtripoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syrtripoli",
		TitleName: "的黎波里",
		TitleCode: "b_syrtripoli",
	}
}
