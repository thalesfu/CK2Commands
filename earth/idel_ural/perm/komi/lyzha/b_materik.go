package lyzha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马捷里克MaterikBarony struct {
	feud.BaseBarony
}

var BMaterik马捷里克 feud.Barony = &马捷里克MaterikBarony{}

func init() {
    f := BMaterik马捷里克.(*马捷里克MaterikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "materik",
		TitleName: "马捷里克",
		TitleCode: "b_materik",
	}
}
