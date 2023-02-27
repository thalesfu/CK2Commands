package prusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿帕米亚ApameaBarony struct {
	feud.BaseBarony
}

var BApamea阿帕米亚 feud.Barony = &阿帕米亚ApameaBarony{}

func init() {
    f := BApamea阿帕米亚.(*阿帕米亚ApameaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apamea",
		TitleName: "阿帕米亚",
		TitleCode: "b_apamea",
	}
}
