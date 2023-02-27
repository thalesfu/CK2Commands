package agen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆瓦萨克MoissacBarony struct {
	feud.BaseBarony
}

var BMoissac穆瓦萨克 feud.Barony = &穆瓦萨克MoissacBarony{}

func init() {
    f := BMoissac穆瓦萨克.(*穆瓦萨克MoissacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moissac",
		TitleName: "穆瓦萨克",
		TitleCode: "b_moissac",
	}
}
