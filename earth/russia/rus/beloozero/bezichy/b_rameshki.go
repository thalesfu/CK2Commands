package bezichy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 季杰里诺RameshkiBarony struct {
	feud.BaseBarony
}

var BRameshki季杰里诺 feud.Barony = &季杰里诺RameshkiBarony{}

func init() {
	f := BRameshki季杰里诺.(*季杰里诺RameshkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rameshki",
		TitleName: "季杰里诺",
		TitleCode: "b_rameshki",
	}
}
