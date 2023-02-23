package chalkidike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波利伊罗斯PolygyrosBarony struct {
	feud.BaseBarony
}

var BPolygyros波利伊罗斯 feud.Barony = &波利伊罗斯PolygyrosBarony{}

func init() {
	f := BPolygyros波利伊罗斯.(*波利伊罗斯PolygyrosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "polygyros",
		TitleName: "波利伊罗斯",
		TitleCode: "b_polygyros",
	}
}
