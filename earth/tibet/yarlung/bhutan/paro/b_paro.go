package paro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕罗ParoBarony struct {
	feud.BaseBarony
}

var BParo帕罗 feud.Barony = &帕罗ParoBarony{}

func init() {
    f := BParo帕罗.(*帕罗ParoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paro",
		TitleName: "帕罗",
		TitleCode: "b_paro",
	}
}
