package poher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔莱CorlayBarony struct {
	feud.BaseBarony
}

var BCorlay科尔莱 feud.Barony = &科尔莱CorlayBarony{}

func init() {
    f := BCorlay科尔莱.(*科尔莱CorlayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corlay",
		TitleName: "科尔莱",
		TitleCode: "b_corlay",
	}
}
