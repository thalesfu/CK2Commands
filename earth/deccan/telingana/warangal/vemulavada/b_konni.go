package vemulavada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘尼KonniBarony struct {
	feud.BaseBarony
}

var BKonni拘尼 feud.Barony = &拘尼KonniBarony{}

func init() {
    f := BKonni拘尼.(*拘尼KonniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konni",
		TitleName: "拘尼",
		TitleCode: "b_konni",
	}
}
