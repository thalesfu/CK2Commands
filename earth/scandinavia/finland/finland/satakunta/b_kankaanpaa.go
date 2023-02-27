package satakunta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎康佩KankaanpaaBarony struct {
	feud.BaseBarony
}

var BKankaanpaa坎康佩 feud.Barony = &坎康佩KankaanpaaBarony{}

func init() {
    f := BKankaanpaa坎康佩.(*坎康佩KankaanpaaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kankaanpaa",
		TitleName: "坎康佩",
		TitleCode: "b_kankaanpaa",
	}
}
