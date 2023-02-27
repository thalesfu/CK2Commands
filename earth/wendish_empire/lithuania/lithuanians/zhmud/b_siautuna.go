package zhmud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希奥图纳SiautunaBarony struct {
	feud.BaseBarony
}

var BSiautuna希奥图纳 feud.Barony = &希奥图纳SiautunaBarony{}

func init() {
    f := BSiautuna希奥图纳.(*希奥图纳SiautunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siautuna",
		TitleName: "希奥图纳",
		TitleCode: "b_siautuna",
	}
}
