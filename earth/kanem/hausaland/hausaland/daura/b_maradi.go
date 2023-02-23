package daura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉迪MaradiBarony struct {
	feud.BaseBarony
}

var BMaradi马拉迪 feud.Barony = &马拉迪MaradiBarony{}

func init() {
	f := BMaradi马拉迪.(*马拉迪MaradiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maradi",
		TitleName: "马拉迪",
		TitleCode: "b_maradi",
	}
}
