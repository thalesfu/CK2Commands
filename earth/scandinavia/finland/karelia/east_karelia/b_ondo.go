package east_karelia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 翁多OndoBarony struct {
	feud.BaseBarony
}

var BOndo翁多 feud.Barony = &翁多OndoBarony{}

func init() {
    f := BOndo翁多.(*翁多OndoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ondo",
		TitleName: "翁多",
		TitleCode: "b_ondo",
	}
}
