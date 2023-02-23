package baalbek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里亚格RiyaqBarony struct {
	feud.BaseBarony
}

var BRiyaq里亚格 feud.Barony = &里亚格RiyaqBarony{}

func init() {
	f := BRiyaq里亚格.(*里亚格RiyaqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "riyaq",
		TitleName: "里亚格",
		TitleCode: "b_riyaq",
	}
}
