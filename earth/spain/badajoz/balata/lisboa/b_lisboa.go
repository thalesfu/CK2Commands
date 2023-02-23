package lisboa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里斯本LisboaBarony struct {
	feud.BaseBarony
}

var BLisboa里斯本 feud.Barony = &里斯本LisboaBarony{}

func init() {
	f := BLisboa里斯本.(*里斯本LisboaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lisboa",
		TitleName: "里斯本",
		TitleCode: "b_lisboa",
	}
}
