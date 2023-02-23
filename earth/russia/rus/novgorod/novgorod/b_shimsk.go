package novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希姆斯克ShimskBarony struct {
	feud.BaseBarony
}

var BShimsk希姆斯克 feud.Barony = &希姆斯克ShimskBarony{}

func init() {
	f := BShimsk希姆斯克.(*希姆斯克ShimskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shimsk",
		TitleName: "希姆斯克",
		TitleCode: "b_shimsk",
	}
}
