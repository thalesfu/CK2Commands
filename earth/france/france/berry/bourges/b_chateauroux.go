package bourges

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙托鲁ChateaurouxBarony struct {
	feud.BaseBarony
}

var BChateauroux沙托鲁 feud.Barony = &沙托鲁ChateaurouxBarony{}

func init() {
    f := BChateauroux沙托鲁.(*沙托鲁ChateaurouxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateauroux",
		TitleName: "沙托鲁",
		TitleCode: "b_chateauroux",
	}
}
