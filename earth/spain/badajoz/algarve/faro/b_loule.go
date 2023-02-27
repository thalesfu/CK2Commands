package faro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛莱LouleBarony struct {
	feud.BaseBarony
}

var BLoule洛莱 feud.Barony = &洛莱LouleBarony{}

func init() {
    f := BLoule洛莱.(*洛莱LouleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loule",
		TitleName: "洛莱",
		TitleCode: "b_loule",
	}
}
