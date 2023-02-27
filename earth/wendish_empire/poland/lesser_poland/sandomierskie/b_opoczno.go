package sandomierskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥波奇诺OpocznoBarony struct {
	feud.BaseBarony
}

var BOpoczno奥波奇诺 feud.Barony = &奥波奇诺OpocznoBarony{}

func init() {
    f := BOpoczno奥波奇诺.(*奥波奇诺OpocznoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "opoczno",
		TitleName: "奥波奇诺",
		TitleCode: "b_opoczno",
	}
}
