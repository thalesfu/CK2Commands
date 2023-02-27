package yazd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔安吉尔Dar_anjirBarony struct {
	feud.BaseBarony
}

var BDar_anjir达尔安吉尔 feud.Barony = &达尔安吉尔Dar_anjirBarony{}

func init() {
    f := BDar_anjir达尔安吉尔.(*达尔安吉尔Dar_anjirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dar_anjir",
		TitleName: "达尔安吉尔",
		TitleCode: "b_dar_anjir",
	}
}
