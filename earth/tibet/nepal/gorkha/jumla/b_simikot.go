package jumla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯弥拘吒SimikotBarony struct {
	feud.BaseBarony
}

var BSimikot斯弥拘吒 feud.Barony = &斯弥拘吒SimikotBarony{}

func init() {
    f := BSimikot斯弥拘吒.(*斯弥拘吒SimikotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "simikot",
		TitleName: "斯弥拘吒",
		TitleCode: "b_simikot",
	}
}
