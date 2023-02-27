package kyunglung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顿久DolchuBarony struct {
	feud.BaseBarony
}

var BDolchu顿久 feud.Barony = &顿久DolchuBarony{}

func init() {
    f := BDolchu顿久.(*顿久DolchuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dolchu",
		TitleName: "顿久",
		TitleCode: "b_dolchu",
	}
}
