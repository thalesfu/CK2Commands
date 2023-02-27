package novogrodek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 休钦ShchuchynBarony struct {
	feud.BaseBarony
}

var BShchuchyn休钦 feud.Barony = &休钦ShchuchynBarony{}

func init() {
    f := BShchuchyn休钦.(*休钦ShchuchynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shchuchyn",
		TitleName: "休钦",
		TitleCode: "b_shchuchyn",
	}
}
