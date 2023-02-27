package innse_gall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓韦根DunveganBarony struct {
	feud.BaseBarony
}

var BDunvegan邓韦根 feud.Barony = &邓韦根DunveganBarony{}

func init() {
    f := BDunvegan邓韦根.(*邓韦根DunveganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunvegan",
		TitleName: "邓韦根",
		TitleCode: "b_dunvegan",
	}
}
