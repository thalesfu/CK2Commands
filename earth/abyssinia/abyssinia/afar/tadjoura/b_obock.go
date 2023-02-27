package tadjoura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥博克ObockBarony struct {
	feud.BaseBarony
}

var BObock奥博克 feud.Barony = &奥博克ObockBarony{}

func init() {
    f := BObock奥博克.(*奥博克ObockBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "obock",
		TitleName: "奥博克",
		TitleCode: "b_obock",
	}
}
