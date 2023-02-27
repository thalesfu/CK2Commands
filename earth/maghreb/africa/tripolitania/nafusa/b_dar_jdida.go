package nafusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔吉达Dar_jdidaBarony struct {
	feud.BaseBarony
}

var BDar_jdida达尔吉达 feud.Barony = &达尔吉达Dar_jdidaBarony{}

func init() {
    f := BDar_jdida达尔吉达.(*达尔吉达Dar_jdidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dar_jdida",
		TitleName: "达尔吉达",
		TitleCode: "b_dar_jdida",
	}
}
