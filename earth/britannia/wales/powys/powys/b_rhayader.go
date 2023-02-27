package powys

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里达尔特RhayaderBarony struct {
	feud.BaseBarony
}

var BRhayader里达尔特 feud.Barony = &里达尔特RhayaderBarony{}

func init() {
    f := BRhayader里达尔特.(*里达尔特RhayaderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rhayader",
		TitleName: "里达尔特",
		TitleCode: "b_rhayader",
	}
}
