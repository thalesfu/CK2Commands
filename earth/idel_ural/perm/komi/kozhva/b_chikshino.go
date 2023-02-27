package kozhva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇克希诺ChikshinoBarony struct {
	feud.BaseBarony
}

var BChikshino奇克希诺 feud.Barony = &奇克希诺ChikshinoBarony{}

func init() {
    f := BChikshino奇克希诺.(*奇克希诺ChikshinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chikshino",
		TitleName: "奇克希诺",
		TitleCode: "b_chikshino",
	}
}
