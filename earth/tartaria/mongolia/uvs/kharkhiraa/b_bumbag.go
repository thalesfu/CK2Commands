package kharkhiraa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奔巴图BumbagBarony struct {
	feud.BaseBarony
}

var BBumbag奔巴图 feud.Barony = &奔巴图BumbagBarony{}

func init() {
    f := BBumbag奔巴图.(*奔巴图BumbagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bumbag",
		TitleName: "奔巴图",
		TitleCode: "b_bumbag",
	}
}
