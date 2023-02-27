package lettigalians

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔梅尔WolmerBarony struct {
	feud.BaseBarony
}

var BWolmer沃尔梅尔 feud.Barony = &沃尔梅尔WolmerBarony{}

func init() {
    f := BWolmer沃尔梅尔.(*沃尔梅尔WolmerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wolmer",
		TitleName: "沃尔梅尔",
		TitleCode: "b_wolmer",
	}
}
