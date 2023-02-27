package kodalaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提婆姞利呬DeogarhBarony struct {
	feud.BaseBarony
}

var BDeogarh提婆姞利呬 feud.Barony = &提婆姞利呬DeogarhBarony{}

func init() {
    f := BDeogarh提婆姞利呬.(*提婆姞利呬DeogarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deogarh",
		TitleName: "提婆姞利呬",
		TitleCode: "b_deogarh",
	}
}
