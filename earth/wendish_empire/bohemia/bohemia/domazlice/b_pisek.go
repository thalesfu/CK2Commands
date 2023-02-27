package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮塞克PisekBarony struct {
	feud.BaseBarony
}

var BPisek皮塞克 feud.Barony = &皮塞克PisekBarony{}

func init() {
    f := BPisek皮塞克.(*皮塞克PisekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pisek",
		TitleName: "皮塞克",
		TitleCode: "b_pisek",
	}
}
