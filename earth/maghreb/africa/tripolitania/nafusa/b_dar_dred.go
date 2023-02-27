package nafusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达尔德雷德Dar_dredBarony struct {
	feud.BaseBarony
}

var BDar_dred达尔德雷德 feud.Barony = &达尔德雷德Dar_dredBarony{}

func init() {
    f := BDar_dred达尔德雷德.(*达尔德雷德Dar_dredBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dar_dred",
		TitleName: "达尔德雷德",
		TitleCode: "b_dar_dred",
	}
}
