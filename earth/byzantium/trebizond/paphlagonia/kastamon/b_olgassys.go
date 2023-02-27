package kastamon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔加西斯OlgassysBarony struct {
	feud.BaseBarony
}

var BOlgassys奥尔加西斯 feud.Barony = &奥尔加西斯OlgassysBarony{}

func init() {
    f := BOlgassys奥尔加西斯.(*奥尔加西斯OlgassysBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olgassys",
		TitleName: "奥尔加西斯",
		TitleCode: "b_olgassys",
	}
}
