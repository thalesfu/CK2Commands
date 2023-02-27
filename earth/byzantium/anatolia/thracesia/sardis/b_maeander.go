package sardis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迈安德MaeanderBarony struct {
	feud.BaseBarony
}

var BMaeander迈安德 feud.Barony = &迈安德MaeanderBarony{}

func init() {
    f := BMaeander迈安德.(*迈安德MaeanderBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maeander",
		TitleName: "迈安德",
		TitleCode: "b_maeander",
	}
}
