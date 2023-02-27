package cornwall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里斯托默尔RestormelBarony struct {
	feud.BaseBarony
}

var BRestormel里斯托默尔 feud.Barony = &里斯托默尔RestormelBarony{}

func init() {
    f := BRestormel里斯托默尔.(*里斯托默尔RestormelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "restormel",
		TitleName: "里斯托默尔",
		TitleCode: "b_restormel",
	}
}
