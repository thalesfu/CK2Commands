package wolgast

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔加斯特WolgastBarony struct {
	feud.BaseBarony
}

var BWolgast沃尔加斯特 feud.Barony = &沃尔加斯特WolgastBarony{}

func init() {
    f := BWolgast沃尔加斯特.(*沃尔加斯特WolgastBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wolgast",
		TitleName: "沃尔加斯特",
		TitleCode: "b_wolgast",
	}
}
