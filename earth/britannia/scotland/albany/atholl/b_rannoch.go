package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰诺赫RannochBarony struct {
	feud.BaseBarony
}

var BRannoch兰诺赫 feud.Barony = &兰诺赫RannochBarony{}

func init() {
    f := BRannoch兰诺赫.(*兰诺赫RannochBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rannoch",
		TitleName: "兰诺赫",
		TitleCode: "b_rannoch",
	}
}
