package astorga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里瓦德拉戈RibadelagoBarony struct {
	feud.BaseBarony
}

var BRibadelago里瓦德拉戈 feud.Barony = &里瓦德拉戈RibadelagoBarony{}

func init() {
    f := BRibadelago里瓦德拉戈.(*里瓦德拉戈RibadelagoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ribadelago",
		TitleName: "里瓦德拉戈",
		TitleCode: "b_ribadelago",
	}
}
