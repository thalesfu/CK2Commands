package mandavyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那菟罗NaddulaBarony struct {
	feud.BaseBarony
}

var BNaddula那菟罗 feud.Barony = &那菟罗NaddulaBarony{}

func init() {
	f := BNaddula那菟罗.(*那菟罗NaddulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naddula",
		TitleName: "那菟罗",
		TitleCode: "b_naddula",
	}
}
